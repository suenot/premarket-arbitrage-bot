# Premarket Arbitrage Bot

Бот для мониторинга и исполнения арбитражных возможностей на prediction markets через [Premarket.me](https://premarket.me) API.

## Быстрый старт

```bash
# 1. Настроить .env
cp .env.example .env
# Заполнить PREMARKET_API_KEY, ETH_WALLET, ETH_SECRET

# 2. Собрать
go build -o arbitrage-bot .

# 3. Запустить (dry-run по умолчанию)
./arbitrage-bot
```

## Документация

| Документ | Описание |
|----------|----------|
| [docs/architecture.md](docs/architecture.md) | Архитектура и структура проекта |
| [docs/api.md](docs/api.md) | Premarket.me API: эндпоинты, схемы, примеры |
| [docs/platforms.md](docs/platforms.md) | Платформы prediction markets |
| [docs/configuration.md](docs/configuration.md) | Все параметры конфигурации |
| [docs/arbitrage-logic.md](docs/arbitrage-logic.md) | Логика арбитража: как это работает |
| [docs/execution.md](docs/execution.md) | Polymarket CLOB execution (EIP-712) |

## Фильтры по умолчанию

| Параметр | Значение | Описание |
|----------|----------|----------|
| `MIN_PROFIT_PCT` | 3% | Минимальный profit |
| `MIN_APR` | 10% | Минимальная годовая доходность |
| `MIN_DEPTH_USD` | $10 | Минимальная ликвидность |
| `DRY_RUN` | true | Только мониторинг, без торговли |

## Лицензия

MIT
