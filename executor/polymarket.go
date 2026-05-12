package executor

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/suenot/w_premarket_arbitrage/scanner"
)

// PolymarketExecutor handles order placement on Polymarket.
// Currently operates as a logger for dry-run mode.
// Full CLOB integration (EIP-712 + HMAC auth) will be added when live trading is enabled.
type PolymarketExecutor struct {
	privateKey *ecdsa.PrivateKey
	wallet     string
	dryRun     bool
	maxUSD     float64
}

// New creates a new PolymarketExecutor.
func New(privKey *ecdsa.PrivateKey, wallet string, dryRun bool, maxUSD float64) *PolymarketExecutor {
	return &PolymarketExecutor{
		privateKey: privKey,
		wallet:     wallet,
		dryRun:     dryRun,
		maxUSD:     maxUSD,
	}
}

// CanExecute checks if both sides of the arbitrage can be executed via Polymarket.
func (e *PolymarketExecutor) CanExecute(pair scanner.ArbitragePair) bool {
	return pair.Arbitrage.BuyYesSource == "polymarket" || pair.Arbitrage.BuyNoSource == "polymarket"
}

// IsFullyExecutable checks if BOTH sides are on Polymarket (full auto-execution possible).
func (e *PolymarketExecutor) IsFullyExecutable(pair scanner.ArbitragePair) bool {
	return pair.Arbitrage.BuyYesSource == "polymarket" && pair.Arbitrage.BuyNoSource == "polymarket"
}

// Execute attempts to execute an arbitrage opportunity.
func (e *PolymarketExecutor) Execute(pair scanner.ArbitragePair) error {
	if e.dryRun {
		log.Printf("[DRY RUN] Would execute arbitrage: %s", pair.Platform1.Question)
		log.Printf("[DRY RUN]   Buy YES @ %s: %.3f (token: %s)",
			pair.Arbitrage.BuyYesSource, pair.Arbitrage.BuyYesPrice, getYesTokenID(pair))
		log.Printf("[DRY RUN]   Buy NO  @ %s: %.3f (token: %s)",
			pair.Arbitrage.BuyNoSource, pair.Arbitrage.BuyNoPrice, getNoTokenID(pair))
		log.Printf("[DRY RUN]   Total cost: %.4f, Profit: %.4f (%.1f%%), APR: %.1f%%",
			pair.Arbitrage.TotalCost, pair.Arbitrage.Profit, pair.Arbitrage.ProfitPct, pair.Arbitrage.APR)
		log.Printf("[DRY RUN]   Max depth: $%.2f, Trade limit: $%.2f",
			pair.Depth.MaxUSD, e.maxUSD)
		return nil
	}

	// Check depth
	tradeSize := pair.Depth.MaxUSD
	if tradeSize > e.maxUSD {
		tradeSize = e.maxUSD
	}
	if tradeSize < 1.0 {
		return fmt.Errorf("insufficient depth: $%.2f", tradeSize)
	}

	// TODO: Implement actual Polymarket CLOB order placement:
	// 1. Derive API credentials via L1 EIP-712 auth
	// 2. Place BUY YES order on buy_yes_source
	// 3. Place BUY NO order on buy_no_source
	// 4. Monitor fills
	//
	// Requires:
	// - EIP-712 typed data signing for Polymarket CTF Exchange
	// - HMAC headers for L2 API authentication
	// - USDC approval on Polygon network
	//
	// See: https://docs.polymarket.com/

	return fmt.Errorf("live execution not yet implemented — use DRY_RUN=true")
}

func getYesTokenID(pair scanner.ArbitragePair) string {
	if pair.Arbitrage.BuyYesSource == pair.Platform1.Source {
		return pair.Platform1.YesTokenID
	}
	return pair.Platform2.YesTokenID
}

func getNoTokenID(pair scanner.ArbitragePair) string {
	if pair.Arbitrage.BuyNoSource == pair.Platform1.Source {
		return pair.Platform1.NoTokenID
	}
	return pair.Platform2.NoTokenID
}
