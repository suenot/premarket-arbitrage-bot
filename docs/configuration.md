# Конфигурация

Все параметры настраиваются через `.env` файл в корне проекта.

## Шаблон

```env
# === Кошелёк (обязательно) ===
ETH_WALLET=0xYourWalletAddress
ETH_SECRET="twelve word mnemonic seed phrase from metamask wallet here"

# === Настройки бота ===
DRY_RUN=true
MIN_PROFIT_PCT=3
MIN_APR=10
MIN_DEPTH_USD=10
POLL_INTERVAL=30
MAX_TRADE_USD=50
```

## Параметры

### Аутентификация

| Параметр | Тип | Описание |
|----------|-----|----------|
| `ETH_WALLET` | string | Адрес Ethereum-кошелька (0x...) для верификации |
| `ETH_SECRET` | string | **Обязательно.** 12-словная BIP-39 мнемоника из MetaMask |

**Авторизация**: бот автоматически подписывает сообщение приватным ключом → получает fresh token от Premarket API.

### Настройки бота

| Параметр | Тип | По умолчанию | Описание |
|----------|-----|-------------|----------|
| `DRY_RUN` | bool | `true` | `true` = мониторинг, `false` = торговля |
| `MIN_PROFIT_PCT` | float | `3.0` | Минимальный profit % |
| `MIN_APR` | float | `10.0` | Минимальная годовая доходность |
| `MIN_DEPTH_USD` | float | `10.0` | Минимальная ликвидность в USD |
| `POLL_INTERVAL` | int | `30` | Интервал опроса API (секунды) |
| `MAX_TRADE_USD` | float | `50.0` | Макс. сумма одной сделки |

## Профили

### Консервативный (default)
```env
MIN_PROFIT_PCT=3
MIN_APR=10
MIN_DEPTH_USD=10
MAX_TRADE_USD=50
```

### Агрессивный
```env
MIN_PROFIT_PCT=1
MIN_APR=5
MIN_DEPTH_USD=5
MAX_TRADE_USD=200
```

## Безопасность

- `.env` и `.env.*` — в `.gitignore` (кроме `.env.example`)
- `temp/` — в `.gitignore`
- Seed phrase даёт полный доступ к кошельку — никогда не коммитить

## Деривация кошелька

Путь: `m/44'/60'/0'/0/0` (стандарт MetaMask).

При запуске бот:
1. Деривирует ECDSA private key из мнемоники
2. Верифицирует адрес (если `ETH_WALLET` указан)
3. Подписывает auth-сообщение для Premarket API
4. Получает fresh access_token + refresh_token

```
🔑 Deriving wallet from mnemonic...
✅ Wallet verified: 0x1F512378f7559DD0581eF343FA0f7244F382D5af
🔐 Premarket auth successful (access token received)
```
