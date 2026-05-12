package premarket

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
)

// AuthResponse is returned by sign-in and contains session tokens.
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// APIKeyResponse is returned by API key creation.
type APIKeyResponse struct {
	APIKey string `json:"api_key"`
}

// SignInWithWallet authenticates using an Ethereum wallet signature.
func (c *Client) SignInWithWallet(privateKey *ecdsa.PrivateKey, chainID int) error {
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Sign the auth message (EIP-191 personal_sign)
	message := fmt.Sprintf("Sign in to Premarket with address %s", address.Hex())
	hash := signPersonalMessage([]byte(message))
	sig, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return fmt.Errorf("signing failed: %w", err)
	}

	// Adjust V value for Ethereum compatibility (27/28)
	if sig[64] < 27 {
		sig[64] += 27
	}

	sigHex := fmt.Sprintf("0x%x", sig)

	body := map[string]interface{}{
		"address":   address.Hex(),
		"chain_id":  chainID,
		"signature": sigHex,
	}

	respBody, err := c.postNoAuth("/api/auth", body)
	if err != nil {
		return fmt.Errorf("sign-in request failed: %w", err)
	}

	var authResp AuthResponse
	if err := json.Unmarshal(respBody, &authResp); err != nil {
		return fmt.Errorf("parse auth response: %w (body: %s)", err, string(respBody))
	}

	if authResp.AccessToken == "" {
		return fmt.Errorf("sign-in returned empty access token (body: %s)", string(respBody))
	}

	c.SetTokens(authResp.AccessToken, authResp.RefreshToken)
	log.Printf("🔐 Premarket auth successful (access token received)")

	// Set up re-auth function for automatic token renewal
	c.SetReAuthFn(func() error {
		log.Println("🔄 Re-authenticating with Premarket...")
		return c.SignInWithWallet(privateKey, chainID)
	})

	return nil
}

// CreateAPIKey generates a long-lived API key for the authenticated user.
func (c *Client) CreateAPIKey() (string, error) {
	respBody, err := c.Post("/api/api-key/generate", nil)
	if err != nil {
		return "", fmt.Errorf("create API key: %w", err)
	}

	var resp APIKeyResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return "", fmt.Errorf("parse API key response: %w", err)
	}

	if resp.APIKey != "" {
		c.SetAPIKey(resp.APIKey)
	}

	return resp.APIKey, nil
}

// postNoAuth performs a POST without Authorization header (for sign-in).
func (c *Client) postNoAuth(path string, body interface{}) ([]byte, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("POST %s returned %d: %s", path, resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// signPersonalMessage creates the Ethereum personal sign hash.
func signPersonalMessage(message []byte) []byte {
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))
	return crypto.Keccak256(append([]byte(prefix), message...))
}
