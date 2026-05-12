package premarket

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// PlatformOrderbook represents orderbook data for a single platform.
type PlatformOrderbook struct {
	YesAsk   float64     `json:"yes_ask"`
	YesBid   float64     `json:"yes_bid"`
	NoAsk    float64     `json:"no_ask"`
	NoBid    float64     `json:"no_bid"`
	Yes      OrderSides  `json:"yes"`
	No       OrderSides  `json:"no"`
	YesPrice float64     `json:"yes_price"`
	NoPrice  float64     `json:"no_price"`
}

// OrderSides contains ask and bid levels.
type OrderSides struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

// OrderbookResponse contains orderbooks from all platforms.
type OrderbookResponse struct {
	Polymarket *PlatformOrderbook `json:"polymarket,omitempty"`
	Opinion    *PlatformOrderbook `json:"opinion,omitempty"`
	Probable   *PlatformOrderbook `json:"probable,omitempty"`
	PredictFun *PlatformOrderbook `json:"predictfun,omitempty"`
	Limitless  *PlatformOrderbook `json:"limitless,omitempty"`
	Combined   *PlatformOrderbook `json:"combined,omitempty"`
}

// OrderbookParams specifies token IDs for the orderbook query.
type OrderbookParams struct {
	PMYesToken   string // Polymarket YES token id
	PMNoToken    string // Polymarket NO token id
	OpYesToken   string // Opinion YES token id
	OpNoToken    string // Opinion NO token id
	ProbYesToken string // Probable YES token id
	ProbNoToken  string // Probable NO token id
	PFMarketID   string // Predict.fun market id
	LimMarketSlug string // Limitless market slug
	LimYesToken  string // Limitless YES token id
	LimNoToken   string // Limitless NO token id
}

// GetOrderbook returns price levels for one outcome across all platforms.
func (c *Client) GetOrderbook(params OrderbookParams) (*OrderbookResponse, error) {
	q := url.Values{}
	if params.PMYesToken != "" {
		q.Set("pm_yes_token", params.PMYesToken)
	}
	if params.PMNoToken != "" {
		q.Set("pm_no_token", params.PMNoToken)
	}
	if params.OpYesToken != "" {
		q.Set("op_yes_token", params.OpYesToken)
	}
	if params.OpNoToken != "" {
		q.Set("op_no_token", params.OpNoToken)
	}
	if params.ProbYesToken != "" {
		q.Set("prob_yes_token", params.ProbYesToken)
	}
	if params.ProbNoToken != "" {
		q.Set("prob_no_token", params.ProbNoToken)
	}
	if params.PFMarketID != "" {
		q.Set("pf_market_id", params.PFMarketID)
	}
	if params.LimMarketSlug != "" {
		q.Set("lim_market_slug", params.LimMarketSlug)
	}
	if params.LimYesToken != "" {
		q.Set("lim_yes_token", params.LimYesToken)
	}
	if params.LimNoToken != "" {
		q.Set("lim_no_token", params.LimNoToken)
	}

	path := "/api/orderbook"
	if encoded := q.Encode(); encoded != "" {
		path += "?" + encoded
	}

	body, err := c.Get(path)
	if err != nil {
		return nil, fmt.Errorf("get orderbook: %w", err)
	}

	var resp OrderbookResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parse orderbook: %w", err)
	}

	return &resp, nil
}

// GetPlatformBook returns the orderbook for a specific platform, or nil.
func (r *OrderbookResponse) GetPlatformBook(source string) *PlatformOrderbook {
	switch source {
	case "polymarket":
		return r.Polymarket
	case "opinion":
		return r.Opinion
	case "probable":
		return r.Probable
	case "predictfun":
		return r.PredictFun
	case "limitless":
		return r.Limitless
	default:
		return nil
	}
}
