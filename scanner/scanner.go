package scanner

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/suenot/w_premarket_arbitrage/premarket"
)

// PlatformData represents one side of an arbitrage pair.
type PlatformData struct {
	Source     string  `json:"source"`
	MarketID  string  `json:"market_id"`
	EventID   string  `json:"event_id"`
	EventSlug string  `json:"event_slug"`
	EventTitle string `json:"event_title"`
	Outcome   string  `json:"outcome"`
	Question  string  `json:"question"`
	URL       string  `json:"url"`
	YesPrice  float64 `json:"yes_price"`
	NoPrice   float64 `json:"no_price"`
	YesTokenID string `json:"yes_token_id"`
	NoTokenID  string `json:"no_token_id"`
}

// ArbitrageData contains the calculated arbitrage metrics.
type ArbitrageData struct {
	BuyYesPrice  float64 `json:"buy_yes_price"`
	BuyYesSource string  `json:"buy_yes_source"`
	BuyNoPrice   float64 `json:"buy_no_price"`
	BuyNoSource  string  `json:"buy_no_source"`
	TotalCost    float64 `json:"total_cost"`
	Profit       float64 `json:"profit"`
	ProfitPct    float64 `json:"profit_pct"`
	DaysUntil    int     `json:"days_until"`
	APR          float64 `json:"apr"`
	PlatformFee  float64 `json:"platform_fee"`
}

// DepthData contains liquidity information.
type DepthData struct {
	MaxShares float64 `json:"max_shares"`
	MaxUSD    float64 `json:"max_usd"`
	YesUSD    float64 `json:"yes_usd"`
	NoUSD     float64 `json:"no_usd"`
	AvgYes    float64 `json:"avg_yes"`
	AvgNo     float64 `json:"avg_no"`
}

// ArbitragePair represents a single arbitrage opportunity.
type ArbitragePair struct {
	Platform1 PlatformData  `json:"platform1"`
	Platform2 PlatformData  `json:"platform2"`
	Arbitrage ArbitrageData `json:"arbitrage"`
	Depth     DepthData     `json:"depth"`
	EndDate   string        `json:"end_date"`
}

// Response is the top-level API response.
type Response struct {
	Pairs []ArbitragePair `json:"pairs"`
	Total int             `json:"total"`
	Count int             `json:"count"`
}

// Scanner polls the Premarket API for arbitrage opportunities.
type Scanner struct {
	client *premarket.Client
}

// New creates a new Scanner using the Premarket client.
func New(client *premarket.Client) *Scanner {
	return &Scanner{client: client}
}

// Fetch retrieves all arbitrage pairs from the API.
func (s *Scanner) Fetch() (*Response, error) {
	body, err := s.client.Get("/api/arbitrage")
	if err != nil {
		return nil, fmt.Errorf("fetch arbitrage: %w", err)
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parse arbitrage: %w", err)
	}

	return &result, nil
}

// FilterOptions defines the criteria for filtering arbitrage pairs.
type FilterOptions struct {
	MinProfitPct float64
	MinAPR       float64
	MinDepthUSD  float64
}

// Filter returns pairs that match the given criteria, sorted by APR descending.
func Filter(pairs []ArbitragePair, opts FilterOptions) []ArbitragePair {
	var filtered []ArbitragePair
	for _, p := range pairs {
		if p.Arbitrage.ProfitPct >= opts.MinProfitPct &&
			p.Arbitrage.APR >= opts.MinAPR &&
			p.Depth.MaxUSD >= opts.MinDepthUSD {
			filtered = append(filtered, p)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Arbitrage.APR > filtered[j].Arbitrage.APR
	})

	return filtered
}
