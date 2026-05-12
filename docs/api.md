# Premarket.me API

Полная документация: [docs.premarket.me](https://docs.premarket.me/)

Скачанная копия всех страниц: `./temp/` (не коммитится)

## Base URL

```
https://api.premarket.me
```

## Аутентификация

### Способ 1: Wallet Sign-In (рекомендуется для ботов)

Бот автоматически аутентифицируется через подпись сообщения приватным ключом.

#### Flow

```
1. Деривируем private key из seed phrase (BIP-39, m/44'/60'/0'/0/0)
2. Подписываем фиксированное сообщение (EIP-191 personal_sign)
3. POST /api/auth → получаем access_token + refresh_token
4. При 401 → POST /api/auth/refresh или повторная подпись
```

#### Сообщение для подписи

> **ВАЖНО**: Это сообщение было reverse-engineered из JavaScript фронтенда Premarket.me
> (chunk `4698-f2603073e10f2314.js`). Оно может измениться — если auth перестанет работать,
> нужно повторно извлечь сообщение из фронтенда.

```
Welcome to Premarket!

Click to sign in. This request will not trigger a blockchain transaction or cost any gas fees.

Wallet: <address_lowercase>
```

**Ключевые детали:**
- Адрес в **нижнем регистре**: `0x1f512378...` (не `0x1F512378...`)
- Подпись: EIP-191 personal_sign (`\x19Ethereum Signed Message:\n<len><message>`)
- V-value: если `sig[64] < 27`, добавить 27 (Ethereum compatibility)
- Формат подписи в запросе: `0x` + hex

#### POST `/api/auth`

```bash
curl -X POST "https://api.premarket.me/api/auth" \
  -H "Content-Type: application/json" \
  -d '{
    "address": "0x1F512378f7559DD0581eF343FA0f7244F382D5af",
    "chain_id": 137,
    "signature": "0x..."
  }'
```

Response:
```json
{
  "access_token": "eyJ...",
  "refresh_token": "eyJ..."
}
```

#### POST `/api/auth/refresh`

```bash
curl -X POST "https://api.premarket.me/api/auth/refresh" \
  -H "Content-Type: application/json" \
  -d '{"refresh_token": "eyJ..."}'
```

Response:
```json
{"access_token": "eyJ..."}
```

#### POST `/api/api-key/generate`

Создаёт долгоживущий API key (требует `Authorization: Bearer <access_token>`):

```bash
curl -X POST "https://api.premarket.me/api/api-key/generate" \
  -H "Authorization: Bearer eyJ..."
```

Response:
```json
{"api_key": "pk_..."}
```

### Способ 2: Ручной JWT

Если wallet auth не работает, можно вручную скопировать JWT из UI premarket.me:

```
Authorization: Bearer <PREMARKET_API_KEY>
```

JWT содержит: `type: ACCESS`, `iss: premarket`, `sub: <uuid>`, `exp: <timestamp>`

---

## Арбитраж

### GET `/api/arbitrage`

Возвращает текущие арбитражные возможности между платформами.

#### Query Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `limit` | integer | Нет | Кол-во пар (по умолчанию — все) |
| `offset` | integer ≥ 0 | Нет | Offset для пагинации |
| `id` | string | Нет | Arbitrage ID для force-include |

#### Пример

```bash
curl -sS "https://api.premarket.me/api/arbitrage?limit=5" \
  -H "Authorization: Bearer $TOKEN"
```

#### Response

```json
{
  "pairs": [ArbitragePair],
  "total": 213,
  "count": 213
}
```

#### ArbitragePair — полная схема

```json
{
  "platform1": {
    "source": "polymarket",
    "market_id": "1234567",
    "event_id": "12345",
    "event_slug": "will-x-happen",
    "event_title": "Will X happen?",
    "outcome": "Yes",
    "question": "Will X happen by 2027?",
    "url": "https://polymarket.com/event/...",
    "yes_price": 0.65,
    "no_price": 0.38,
    "yes_token_id": "12345...6789",
    "no_token_id": "98765...4321"
  },
  "platform2": { ... },
  "arbitrage": {
    "buy_yes_price": 0.50,
    "buy_yes_source": "kalshi",
    "buy_no_price": 0.38,
    "buy_no_source": "polymarket",
    "total_cost": 0.88,
    "profit": 0.12,
    "profit_pct": 13.64,
    "days_until": 200,
    "apr": 24.89,
    "platform_fee": 0.51
  },
  "depth": {
    "max_shares": 100.0,
    "max_usd": 88.0,
    "yes_usd": 50.0,
    "no_usd": 38.0,
    "avg_yes": 0.50,
    "avg_no": 0.38
  },
  "end_date": "2027-01-02T15:00:00"
}
```

---

## Стакан (Orderbook)

### GET `/api/orderbook`

Кросс-платформенный стакан для одного outcome.

#### Query Parameters

| Parameter | Description |
|-----------|-------------|
| `pm_yes_token` | Polymarket YES token id |
| `pm_no_token` | Polymarket NO token id |
| `op_yes_token` | Opinion YES token id |
| `op_no_token` | Opinion NO token id |
| `prob_yes_token` | Probable YES token id |
| `prob_no_token` | Probable NO token id |
| `pf_market_id` | Predict.fun market id |
| `lim_market_slug` | Limitless market slug |
| `lim_yes_token` | Limitless YES token id |
| `lim_no_token` | Limitless NO token id |

#### Response

```json
{
  "polymarket": {
    "yes_ask": 0.65, "yes_bid": 0.63,
    "no_ask": 0.38, "no_bid": 0.36,
    "yes": {"asks": [[0.65, 100]], "bids": [[0.63, 50]]},
    "no": {"asks": [[0.38, 200]], "bids": [[0.36, 80]]},
    "yes_price": 0.65, "no_price": 0.38
  },
  "kalshi": { ... },
  "combined": { ... }
}
```

---

## Ликвидность

### POST `/api/liquidity/distribution`

Оптимальный сплит ордера по платформам.

#### Body

```json
{
  "tokens": {
    "polymarket": {"outcome": "yes", "token_id": "123456"},
    "opinion": {"outcome": "yes", "token_id": "987654"},
    "predictfun": {"market_id": "991122", "outcome": "yes", "token_id": "555"}
  },
  "direction": "buy",
  "amount_usd": 100,
  "include_fees": true,
  "enforce_min_orders": true,
  "optimize_gas": false,
  "gas_cost_per_source": 5
}
```

#### Response

```json
{
  "combined": {"platform": "combined", "amount": 100, "price": 0.5271},
  "platforms": {
    "opinion": {
      "platform": "opinion", "used": true, "amount": 60, "price": 0.52,
      "best_price": 0.52, "available_liquidity": 1900,
      "liquidity_exhausted": false
    },
    "predictfun": {
      "platform": "predictfun", "used": true, "amount": 40, "price": 0.535,
      "best_price": 0.531, "available_liquidity": 1200
    }
  },
  "distributions": [
    {"platform": "opinion", "amount": 60, "price": 0.52},
    {"platform": "predictfun", "amount": 40, "price": 0.535}
  ],
  "avg_price": 0.5271
}
```

---

## Балансы

### GET `/api/balances`

Стейблкоин-балансы (USDC) по чейнам.

```bash
curl "https://api.premarket.me/api/balances?addresses=0x1F512378..." \
  -H "Authorization: Bearer $TOKEN"
```

```json
{
  "balances": {
    "polygon": {"balance": 5.70, "symbol": "USDC"},
    "bsc": {"balance": 0, "symbol": "USDT"},
    "base": {"balance": 0, "symbol": "USDC"},
    "total": 5.70
  },
  "addresses": ["0x1F512378..."]
}
```

### GET `/api/balances/tokens`

Market token балансы.

```bash
curl "https://api.premarket.me/api/balances/tokens?addresses=0x...&token_ids=123,456"
```

---

## Позиции

### GET `/api/positions/aggregated`

Портфель позиций для одного или нескольких кошельков.

```bash
curl "https://api.premarket.me/api/positions/aggregated?addresses=0x..." \
  -H "Authorization: Bearer $TOKEN"
```

```json
{
  "summary": {
    "total_value": 1248.41,
    "total_positions": 18,
    "active_positions": 13,
    "claimable_value": 57.22,
    "by_platform": {
      "polymarket": 801.12, "opinion": 312.44,
      "probable": 94.85, "limitless": 40
    },
    "by_chain": {"polygon": 841.12, "bsc": 367.29, "base": 40}
  },
  "positions": [
    {
      "token_id": "...", "event_title": "...", "source": "polymarket",
      "outcome": "Yes", "balance": 100, "current_price": 0.65,
      "value": 65.0, "is_active": true, "claim_available": false
    }
  ]
}
```

---

## Другие эндпоинты

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/events/list` | GET | Список событий |
| `/api/events/{event_slug}` | GET | Детали события |
| `/api/v2/market/{id}` | GET | Рынок с outcomes |
| `/api/v2/platforms` | GET | Список платформ |
| `/api/restricts/check` | GET | Региональная доступность |
| `/api/user/me` | GET | Профиль пользователя |

---

## Error Handling

| Code | Описание |
|------|----------|
| 401 | Токен истёк → refresh или re-auth |
| 4xx | Ошибка запроса / валидации |
| 5xx | Ошибка сервера |
