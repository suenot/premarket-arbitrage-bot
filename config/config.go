package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PremarketAPIKey string
	EthWallet       string
	EthSecret       string // 12-word mnemonic
	DryRun          bool
	MinProfitPct    float64
	MinAPR          float64
	MinDepthUSD     float64
	PollInterval    int // seconds
	MaxTradeUSD     float64
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		PremarketAPIKey: os.Getenv("PREMARKET_API_KEY"),
		EthWallet:       os.Getenv("ETH_WALLET"),
		EthSecret:       os.Getenv("ETH_SECRET"),
		DryRun:          true,
		MinProfitPct:    3.0,
		MinAPR:          10.0,
		MinDepthUSD:     10.0,
		PollInterval:    30,
		MaxTradeUSD:     50.0,
	}

	if v := os.Getenv("DRY_RUN"); v != "" {
		cfg.DryRun, _ = strconv.ParseBool(v)
	}
	if v := os.Getenv("MIN_PROFIT_PCT"); v != "" {
		cfg.MinProfitPct, _ = strconv.ParseFloat(v, 64)
	}
	if v := os.Getenv("MIN_APR"); v != "" {
		cfg.MinAPR, _ = strconv.ParseFloat(v, 64)
	}
	if v := os.Getenv("MIN_DEPTH_USD"); v != "" {
		cfg.MinDepthUSD, _ = strconv.ParseFloat(v, 64)
	}
	if v := os.Getenv("POLL_INTERVAL"); v != "" {
		cfg.PollInterval, _ = strconv.Atoi(v)
	}
	if v := os.Getenv("MAX_TRADE_USD"); v != "" {
		cfg.MaxTradeUSD, _ = strconv.ParseFloat(v, 64)
	}

	return cfg, nil
}
