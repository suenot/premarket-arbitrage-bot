package premarket

import (
	"encoding/json"
	"fmt"
	"strings"
)

// PositionsResponse contains the portfolio view.
type PositionsResponse struct {
	Summary   PositionsSummary `json:"summary"`
	Balances  ChainBalances    `json:"balances"`
	Positions []Position       `json:"positions"`
	Addresses []string         `json:"addresses"`
}

// PositionsSummary is the aggregated portfolio summary.
type PositionsSummary struct {
	TotalValue      float64            `json:"total_value"`
	TotalPositions  int                `json:"total_positions"`
	ActivePositions int                `json:"active_positions"`
	ClaimableValue  float64            `json:"claimable_value"`
	ByPlatform      map[string]float64 `json:"by_platform"`
	ByChain         map[string]float64 `json:"by_chain"`
}

// Position is a single open position.
type Position struct {
	TokenID         string  `json:"token_id"`
	EventTitle      string  `json:"event_title"`
	Source          string  `json:"source"`
	Outcome         string  `json:"outcome"`
	MarketID        int     `json:"market_id"`
	ChainID         string  `json:"chain_id"`
	ContractAddress string  `json:"contract_address"`
	Question        string  `json:"question"`
	EventID         string  `json:"event_id"`
	EventSlug       string  `json:"event_slug"`
	ConditionID     string  `json:"condition_id"`
	BalanceRaw      float64 `json:"balance_raw"`
	Balance         float64 `json:"balance"`
	CurrentPrice    float64 `json:"current_price"`
	Value           float64 `json:"value"`
	IsActive        bool    `json:"is_active"`
	MarketStatus    string  `json:"market_status"`
	ClaimAvailable  bool    `json:"claim_available"`
}

// GetPositions returns open positions for the given wallet addresses.
func (c *Client) GetPositions(addresses []string) (*PositionsResponse, error) {
	path := fmt.Sprintf("/api/positions/aggregated?addresses=%s", strings.Join(addresses, ","))

	body, err := c.Get(path)
	if err != nil {
		return nil, fmt.Errorf("get positions: %w", err)
	}

	var resp PositionsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parse positions: %w", err)
	}

	return &resp, nil
}
