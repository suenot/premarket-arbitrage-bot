package premarket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const BaseURL = "https://api.premarket.me"

// Client is the main Premarket API client with automatic auth management.
type Client struct {
	baseURL      string
	httpClient   *http.Client
	accessToken  string
	refreshToken string
	apiKey       string

	// Auth function to re-authenticate when token expires.
	// Set by auth.go after initial sign-in.
	reAuthFn func() error

	mu sync.RWMutex
}

// NewClient creates a new Premarket API client.
// If accessToken is provided, it's used immediately.
// Call SignInWithWallet() to enable auto-refresh.
func NewClient(accessToken string) *Client {
	return &Client{
		baseURL:     BaseURL,
		accessToken: accessToken,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SetTokens updates the access and refresh tokens.
func (c *Client) SetTokens(access, refresh string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.accessToken = access
	c.refreshToken = refresh
}

// SetAPIKey sets the X-API-Key for endpoints that require it.
func (c *Client) SetAPIKey(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.apiKey = key
}

// SetReAuthFn sets the function to call when re-authentication is needed.
func (c *Client) SetReAuthFn(fn func() error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.reAuthFn = fn
}

// AccessToken returns the current access token.
func (c *Client) AccessToken() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.accessToken
}

// RefreshTokenValue returns the current refresh token.
func (c *Client) RefreshTokenValue() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.refreshToken
}

// Get performs an authenticated GET request.
func (c *Client) Get(path string) ([]byte, error) {
	return c.doRequest("GET", path, nil)
}

// Post performs an authenticated POST request with JSON body.
func (c *Client) Post(path string, body interface{}) ([]byte, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshal body: %w", err)
	}
	return c.doRequest("POST", path, data)
}

func (c *Client) doRequest(method, path string, body []byte) ([]byte, error) {
	resp, err := c.executeRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	// If 401, try to refresh/re-auth and retry once
	if resp.StatusCode == http.StatusUnauthorized {
		resp.Body.Close()

		if err := c.tryReAuth(); err != nil {
			return nil, fmt.Errorf("re-auth failed: %w", err)
		}

		resp, err = c.executeRequest(method, path, body)
		if err != nil {
			return nil, err
		}
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API %s %s returned %d: %s", method, path, resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func (c *Client) executeRequest(method, path string, body []byte) (*http.Response, error) {
	url := c.baseURL + path

	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	c.mu.RLock()
	token := c.accessToken
	apiKey := c.apiKey
	c.mu.RUnlock()

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	if apiKey != "" {
		req.Header.Set("X-API-Key", apiKey)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return c.httpClient.Do(req)
}

func (c *Client) tryReAuth() error {
	c.mu.RLock()
	refresh := c.refreshToken
	reAuth := c.reAuthFn
	c.mu.RUnlock()

	// Try refresh first
	if refresh != "" {
		if err := c.doRefresh(refresh); err == nil {
			return nil
		}
	}

	// Fall back to full re-auth
	if reAuth != nil {
		return reAuth()
	}

	return fmt.Errorf("no refresh token or re-auth function available")
}

// doRefresh calls the refresh endpoint.
func (c *Client) doRefresh(refreshToken string) error {
	body, _ := json.Marshal(map[string]string{
		"refresh_token": refreshToken,
	})

	req, err := http.NewRequest("POST", c.baseURL+"/api/auth/refresh", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("refresh returned %d", resp.StatusCode)
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	c.mu.Lock()
	c.accessToken = result.AccessToken
	c.mu.Unlock()

	return nil
}
