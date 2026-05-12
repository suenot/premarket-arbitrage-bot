# Premarket.me API

Полная документация: [docs.premarket.me](https://docs.premarket.me/)

## Base URL

```
https://api.premarket.me
```

## Аутентификация

Все запросы требуют JWT-токен в заголовке:

```
Authorization: Bearer <PREMARKET_API_KEY>
```

API-ключ — JWT с полями:
- `type`: `ACCESS`
- `iss`: `premarket`
- `sub`: UUID аккаунта
- `exp`: timestamp экспирации

---

## Арбитраж

### GET `/api/arbitrage`

Возвращает текущие арбитражные возможности между платформами.

#### Query Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `limit` | integer | Нет | Кол-во пар (по умолчанию — все) |
| `offset` | integer ≥ 0 | Нет | Offset для пагинации |
| `id` | string | Нет | Arbitrage ID для force-include в guest preview |

#### Пример запроса

```bash
curl -sS "https://api.premarket.me/api/arbitrage?limit=5" \
  -H "Authorization: Bearer $PREMARKET_API_KEY"
```

#### Response

```json
{
  "pairs": [ArbitragePair],
  "total": 213,     // Всего кандидатов ДО фильтра прибыльности
  "count": 213      // Возвращённых пар
}
```

#### ArbitragePair

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
    "description": "Resolution criteria...",
    "url": "https://polymarket.com/event/will-x-happen",
    "yes_price": 0.65,
    "no_price": 0.38,
    "yes_token_id": "12345...6789",
    "no_token_id": "98765...4321"
  },
  "platform2": {
    "source": "kalshi",
    "market_id": "9876543",
    "event_id": "KXEVENT",
    "event_slug": "will-x-happen",
    "event_title": "Will X happen?",
    "outcome": "Yes",
    "question": "Will X happen by 2027?",
    "description": "Resolution criteria...",
    "url": "https://kalshi.com/markets/KXEVENT",
    "yes_price": 0.50,
    "no_price": 0.55,
    "yes_token_id": "...",
    "no_token_id": "..."
  },
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

#### Поля ArbitrageData

| Поле | Тип | Описание |
|------|-----|----------|
| `buy_yes_price` | float | Цена покупки YES (с дешёвой платформы) |
| `buy_yes_source` | string | Платформа для покупки YES |
| `buy_no_price` | float | Цена покупки NO (с другой платформы) |
| `buy_no_source` | string | Платформа для покупки NO |
| `total_cost` | float | YES + NO = total (если < 1.0 — есть арбитраж) |
| `profit` | float | `1.0 - total_cost` = гарантированная прибыль на $1 |
| `profit_pct` | float | Profit в процентах |
| `days_until` | int | Дней до экспирации рынка |
| `apr` | float | Годовая доходность (annualized) |
| `platform_fee` | float | Комиссия платформы (%) |

#### Поля DepthData

| Поле | Тип | Описание |
|------|-----|----------|
| `max_shares` | float | Макс. кол-во акций по текущим ценам |
| `max_usd` | float | Макс. объём в USD |
| `yes_usd` | float | Ликвидность YES стороны в USD |
| `no_usd` | float | Ликвидность NO стороны в USD |
| `avg_yes` | float | Средневзвешенная цена YES |
| `avg_no` | float | Средневзвешенная цена NO |

---

## Другие эндпоинты

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/events/list` | GET | Список событий (offset, limit, sort, platform, category) |
| `/api/events/{event_slug}` | GET | Детали события |
| `/api/v2/market/{id}` | GET | Рынок с outcomes |
| `/api/orderbook/{...}` | GET | Стакан по outcome |
| `/api/liquidity/plan` | POST | Оптимальный сплит ордера |
| `/api/positions` | GET | Открытые позиции |
| `/api/balances/tokens` | GET | Балансы токенов |
| `/api/balances/cash` | GET | Балансы стейблов |
| `/api/v2/platforms` | GET | Список платформ |
| `/api/restricts/check` | GET | Региональная доступность |

---

## Error Handling

| Code | Описание |
|------|----------|
| 4xx | Ошибка запроса / авторизации / валидации |
| 5xx | Ошибка сервера или upstream |

## Common Response Shape

```json
{
  "count": 10,    // Кол-во возвращённых записей
  "total": 100,   // Всего записей
  "offset": 0,    // Текущий offset
  "limit": 10     // Текущий limit
}
```
