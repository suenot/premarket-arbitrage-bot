# Premarket Arbitrage Bot

[![Go](https://img.shields.io/badge/Go-1.23-00ADD8?logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Automated arbitrage monitoring and execution bot for prediction markets via [Premarket.me](https://premarket.me) API.

**[🇷🇺 Русская версия](README.ru.md)**

## Features

- 🔐 **Auto wallet auth** — signs in with MetaMask seed phrase, auto-refreshes tokens
- 📡 **Cross-platform scanning** — finds arbitrage between Polymarket, Kalshi, Opinion, Probable, Limitless, Predict.fun
- 📊 **Real-time CLI dashboard** — color-coded table with APR, profit %, depth
- 💰 **Balance & positions** — shows USDC balance and open positions via Premarket API
- 🎯 **Smart filtering** — configurable min profit, APR, and liquidity thresholds
- 📈 **Order planning** — optimal order split via Premarket liquidity API
- 🛡️ **Dry-run mode** — monitor only, no trades by default

## Quick Start

```bash
# 1. Clone
git clone https://github.com/suenot/premarket-arbitrage-bot.git
cd premarket-arbitrage-bot

# 2. Configure
cp .env.example .env
# Fill in ETH_WALLET, ETH_SECRET (12-word MetaMask seed phrase)

# 3. Build
go build -o arbitrage-bot .

# 4. Run (dry-run by default)
./arbitrage-bot
```

## How It Works

```
Seed phrase → Private key → Wallet signature → Premarket Auth
                                                      ↓
                                          Scan arbitrage opportunities
                                                      ↓
                                        Verify prices via orderbook API
                                                      ↓
                                       Plan optimal order split (liquidity API)
                                                      ↓
                                           Check USDC balance
                                                      ↓
                                        Execute on Polymarket CLOB
```

## Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `ETH_SECRET` | — | **Required.** 12-word BIP-39 mnemonic from MetaMask |
| `ETH_WALLET` | — | Wallet address for verification |
| `DRY_RUN` | `true` | Monitor only, no trading |
| `MIN_PROFIT_PCT` | `3` | Minimum profit percentage |
| `MIN_APR` | `10` | Minimum annualized return |
| `MIN_DEPTH_USD` | `10` | Minimum liquidity in USD |
| `POLL_INTERVAL` | `30` | API polling interval (seconds) |
| `MAX_TRADE_USD` | `50` | Maximum trade size |

## Documentation

| Document | Description |
|----------|-------------|
| [Architecture](docs/architecture.md) | Project structure and data flow |
| [API Reference](docs/api.md) | Premarket.me API: endpoints, schemas, auth |
| [Platforms](docs/platforms.md) | Supported prediction market platforms |
| [Configuration](docs/configuration.md) | All environment variables |
| [Arbitrage Logic](docs/arbitrage-logic.md) | How arbitrage detection works |
| [Execution](docs/execution.md) | Polymarket CLOB execution (EIP-712) |
| [Premarket API Docs](premarket-docs/) | Downloaded copy of docs.premarket.me |

## Project Structure

```
├── main.go                   # Entry point, poll loop, graceful shutdown
├── config/config.go          # .env loading
├── wallet/wallet.go          # BIP-39 mnemonic → ECDSA key derivation
├── premarket/                # Full Premarket API client
│   ├── client.go             # HTTP client with auto-auth & token refresh
│   ├── auth.go               # Wallet sign-in (EIP-191 personal_sign)
│   ├── orderbook.go          # Cross-platform price verification
│   ├── liquidity.go          # Optimal order split planning
│   ├── balances.go           # USDC & token balances
│   └── positions.go          # Portfolio positions
├── scanner/scanner.go        # Arbitrage opportunity scanner & filter
├── display/display.go        # CLI dashboard (lipgloss)
├── executor/polymarket.go    # Polymarket CLOB order execution
└── docs/                     # Documentation
```

## License

MIT
