package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/suenot/w_premarket_arbitrage/config"
	"github.com/suenot/w_premarket_arbitrage/display"
	"github.com/suenot/w_premarket_arbitrage/executor"
	"github.com/suenot/w_premarket_arbitrage/premarket"
	"github.com/suenot/w_premarket_arbitrage/scanner"
	"github.com/suenot/w_premarket_arbitrage/wallet"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if cfg.EthSecret == "" {
		log.Fatal("ETH_SECRET is required in .env (12-word mnemonic from MetaMask)")
	}

	// Derive wallet from seed phrase
	log.Println("🔑 Deriving wallet from mnemonic...")
	privKey, addr, err := wallet.DeriveKey(cfg.EthSecret)
	if err != nil {
		log.Fatalf("Wallet derivation failed: %v", err)
	}

	if cfg.EthWallet != "" {
		if err := wallet.Verify(cfg.EthSecret, cfg.EthWallet); err != nil {
			log.Fatalf("Wallet verification failed: %v", err)
		}
		log.Printf("✅ Wallet verified: %s", addr.Hex())
	} else {
		log.Printf("📍 Derived wallet: %s", addr.Hex())
	}

	walletAddr := addr.Hex()

	// Authenticate via wallet signature
	pmClient := premarket.NewClient("")
	log.Println("🔐 Authenticating with Premarket...")
	if err := pmClient.SignInWithWallet(privKey, 137); err != nil {
		log.Fatalf("Premarket auth failed: %v", err)
	}

	exec := executor.New(privKey, addr.Hex(), cfg.DryRun, cfg.MaxTradeUSD)

	// Init scanner using shared Premarket client
	s := scanner.New(pmClient)

	filterOpts := scanner.FilterOptions{
		MinProfitPct: cfg.MinProfitPct,
		MinAPR:       cfg.MinAPR,
		MinDepthUSD:  cfg.MinDepthUSD,
	}

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	mode := "DRY RUN"
	if !cfg.DryRun {
		mode = "LIVE"
	}

	fmt.Println()
	log.Printf("🚀 Starting Premarket Arbitrage Bot [%s]", mode)
	log.Printf("   Filters: profit >= %.1f%%, APR >= %.1f%%, depth >= $%.0f",
		cfg.MinProfitPct, cfg.MinAPR, cfg.MinDepthUSD)
	log.Printf("   Poll interval: %ds | Max trade: $%.0f",
		cfg.PollInterval, cfg.MaxTradeUSD)
	fmt.Println()

	// Main loop
	ticker := time.NewTicker(time.Duration(cfg.PollInterval) * time.Second)
	defer ticker.Stop()

	// Run immediately on start
	poll(s, pmClient, filterOpts, cfg.DryRun, exec, walletAddr, cfg.PollInterval)

	for {
		select {
		case <-ticker.C:
			poll(s, pmClient, filterOpts, cfg.DryRun, exec, walletAddr, cfg.PollInterval)
		case <-stop:
			fmt.Println()
			log.Println("👋 Shutting down gracefully...")
			return
		}
	}
}

func poll(s *scanner.Scanner, pmClient *premarket.Client, opts scanner.FilterOptions, dryRun bool, exec *executor.PolymarketExecutor, walletAddr string, pollInterval int) {
	log.Println("📡 Fetching arbitrage opportunities...")

	resp, err := s.Fetch()
	if err != nil {
		log.Printf("❌ Fetch error: %v", err)
		return
	}

	filtered := scanner.Filter(resp.Pairs, opts)

	// Fetch portfolio info via Premarket API
	var portfolio *display.PortfolioInfo
	if walletAddr != "" {
		portfolio = fetchPortfolio(pmClient, walletAddr)
	}

	// Clear screen for fresh display
	fmt.Print("\033[2J\033[H")

	display.Render(filtered, resp.Total, dryRun, portfolio)

	if exec != nil && !dryRun {
		for _, pair := range filtered {
			if exec.CanExecute(pair) {
				if err := exec.Execute(pair); err != nil {
					log.Printf("❌ Execution error: %v", err)
				}
			}
		}
	} else if exec != nil && dryRun {
		polyCount := 0
		for _, pair := range filtered {
			if exec.CanExecute(pair) {
				polyCount++
			}
		}
		if polyCount > 0 {
			log.Printf("💡 %d pairs have Polymarket side (executable with DRY_RUN=false)", polyCount)
		}
	}

	log.Printf("⏳ Next poll in %ds...", pollInterval)
}

func fetchPortfolio(pmClient *premarket.Client, walletAddr string) *display.PortfolioInfo {
	info := &display.PortfolioInfo{}

	// Fetch cash balances
	balances, err := pmClient.GetCashBalances([]string{walletAddr})
	if err != nil {
		log.Printf("⚠️  Balance fetch: %v", err)
	} else {
		info.TotalBalance = balances.Balances.Total
		info.PolygonBalance = balances.Balances.Polygon.Balance
	}

	// Fetch positions
	positions, err := pmClient.GetPositions([]string{walletAddr})
	if err != nil {
		log.Printf("⚠️  Positions fetch: %v", err)
	} else {
		info.ActivePositions = positions.Summary.ActivePositions
		info.TotalValue = positions.Summary.TotalValue
	}

	return info
}
