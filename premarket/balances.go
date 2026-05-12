package premarket

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CashBalancesResponse contains stablecoin balances across chains.
type CashBalancesResponse struct {
	Balances  ChainBalances `json:"balances"`
	Addresses []string      `json:"addresses"`
	Error     string        `json:"error,omitempty"`
}

// ChainBalances contains per-chain balance data.
type ChainBalances struct {
	Polygon ChainBalance `json:"polygon"`
	BSC     ChainBalance `json:"bsc"`
	Base    ChainBalance `json:"base"`
	Total   float64      `json:"total"`
}

// ChainBalance is the balance on a single chain.
type ChainBalance struct {
	Balance   float64            `json:"balance"`
	Symbol    string             `json:"symbol"`
	Addresses map[string]float64 `json:"addresses"`
}

// TokenBalancesResponse contains market token balances.
type TokenBalancesResponse struct {
	Balances []TokenBalance `json:"balances"`
}

// TokenBalance is a single token balance entry.
type TokenBalance struct {
	TokenID  string  `json:"token_id"`
	Balance  float64 `json:"balance"`
	Address  string  `json:"address"`
	Source   string  `json:"source"`
}

// GetCashBalances returns stablecoin balances for the given wallet addresses.
func (c *Client) GetCashBalances(addresses []string) (*CashBalancesResponse, error) {
	path := fmt.Sprintf("/api/balances?addresses=%s", strings.Join(addresses, ","))

	body, err := c.Get(path)
	if err != nil {
		return nil, fmt.Errorf("get cash balances: %w", err)
	}

	var resp CashBalancesResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parse balances: %w", err)
	}

	return &resp, nil
}

// GetTokenBalances returns market token balances for wallets.
func (c *Client) GetTokenBalances(addresses []string, tokenIDs []string) (*TokenBalancesResponse, error) {
	path := fmt.Sprintf("/api/balances/tokens?addresses=%s", strings.Join(addresses, ","))
	if len(tokenIDs) > 0 {
		path += "&token_ids=" + strings.Join(tokenIDs, ",")
	}

	body, err := c.Get(path)
	if err != nil {
		return nil, fmt.Errorf("get token balances: %w", err)
	}

	var resp TokenBalancesResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parse token balances: %w", err)
	}

	return &resp, nil
}
