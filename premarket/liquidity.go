package premarket

import (
	"encoding/json"
	"fmt"
)

// LiquidityRequest is the body for the liquidity distribution endpoint.
type LiquidityRequest struct {
	Tokens          map[string]TokenInfo `json:"tokens"`
	Direction       string               `json:"direction"` // "buy" or "sell"
	AmountUSD       float64              `json:"amount_usd,omitempty"`
	AmountTokens    float64              `json:"amount_tokens,omitempty"`
	IncludeFees     bool                 `json:"include_fees"`
	EnforceMinOrders bool               `json:"enforce_min_orders"`
	OptimizeGas     bool                 `json:"optimize_gas"`
	GasCostPerSource float64            `json:"gas_cost_per_source,omitempty"`
}

// TokenInfo identifies a token on a specific platform.
type TokenInfo struct {
	Outcome  string `json:"outcome"`
	TokenID  string `json:"token_id"`
	MarketID string `json:"market_id,omitempty"`
}

// LiquidityResponse contains the optimal order distribution.
type LiquidityResponse struct {
	Combined      PlatformDistribution            `json:"combined"`
	Platforms     map[string]PlatformDistribution  `json:"platforms"`
	Distributions []Distribution                  `json:"distributions"`
	AvgPrice      float64                         `json:"avg_price"`
	Error         string                          `json:"error,omitempty"`
}

// PlatformDistribution shows allocation for a single platform.
type PlatformDistribution struct {
	Platform           string  `json:"platform"`
	Used               bool    `json:"used"`
	Amount             float64 `json:"amount"`
	Price              float64 `json:"price"`
	BestPrice          float64 `json:"best_price"`
	HypotheticalTokens float64 `json:"hypothetical_tokens"`
	HypotheticalPrice  float64 `json:"hypothetical_price"`
	AvailableLiquidity float64 `json:"available_liquidity"`
	LiquidityExhausted bool    `json:"liquidity_exhausted"`
}

// Distribution is a single allocation entry.
type Distribution struct {
	Platform string  `json:"platform"`
	Amount   float64 `json:"amount"`
	Price    float64 `json:"price"`
}

// PlanOrderSplit calculates the optimal order split across platforms.
func (c *Client) PlanOrderSplit(req LiquidityRequest) (*LiquidityResponse, error) {
	body, err := c.Post("/api/liquidity/distribution", req)
	if err != nil {
		return nil, fmt.Errorf("plan order split: %w", err)
	}

	var resp LiquidityResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parse liquidity response: %w", err)
	}

	if resp.Error != "" {
		return nil, fmt.Errorf("liquidity error: %s", resp.Error)
	}

	return &resp, nil
}
