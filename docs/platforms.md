# Платформы Prediction Markets

Premarket.me агрегирует данные с 6 платформ prediction markets.

## Обзор

| # | Source | Название | Тип | Сеть | Торговля через API | Сайт |
|---|--------|----------|-----|------|-------------------|------|
| 1 | `polymarket` | Polymarket | DEX (on-chain) | Polygon | ✅ CLOB API + SDK | [polymarket.com](https://polymarket.com) |
| 2 | `kalshi` | Kalshi | CEX (регулируемая) | — | ✅ REST API (KYC, US) | [kalshi.com](https://kalshi.com) |
| 3 | `predictfun` | Predict.fun | DEX (on-chain) | Blast | ⚠️ Нет публичного SDK | [predict.fun](https://predict.fun) |
| 4 | `opinion` | Opinion | DEX (on-chain) | — | ⚠️ Нет публичного SDK | — |
| 5 | `probable` | Probable | DEX (on-chain) | — | ⚠️ Нет публичного SDK | — |
| 6 | `limitless` | Limitless | DEX (on-chain) | — | ⚠️ Нет публичного SDK | — |

## Получение списка платформ

```bash
curl -sS "https://api.premarket.me/api/v2/platforms" \
  -H "Authorization: Bearer $TOKEN"
```

```json
{
  "platforms": [
    { "source": "polymarket", "name": "Polymarket" },
    { "source": "kalshi", "name": "Kalshi" },
    { "source": "opinion", "name": "Opinion" },
    { "source": "probable", "name": "Probable" },
    { "source": "predictfun", "name": "Predict.fun" },
    { "source": "limitless", "name": "Limitless" }
  ]
}
```

---

## Polymarket

Крупнейший DEX prediction market на Polygon.

### Ключевые особенности
- **Settlement**: On-chain через CTF (Conditional Token Framework)
- **Торговля**: CLOB (Central Limit Order Book) с off-chain matching
- **Валюта**: USDC на Polygon
- **API**: Полноценный REST + WebSocket
- **SDK**: [py-clob-client](https://github.com/Polymarket/py-clob-client) (Python), community Go SDKs
- **Auth**: EIP-712 signing для L1, HMAC для L2

### Требования для торговли
1. ETH-кошелёк с private key (или seed phrase)
2. USDC на Polygon
3. Approve контрактов Polymarket на spending USDC
4. API credentials (деривируются из wallet signature)

### Go SDK (community)
- [GoPolymarket/polymarket-go-sdk](https://github.com/GoPolymarket/polymarket-go-sdk)
- [0xNetuser/Polymarket-golang](https://github.com/0xNetuser/Polymarket-golang)
- [nijaru/go-clob-client](https://github.com/nijaru/go-clob-client)

### API URLs
- CLOB API: `https://clob.polymarket.com`
- Gamma API (metadata): `https://gamma-api.polymarket.com`
- Chain: Polygon (chain_id: 137)

---

## Kalshi

Регулируемая (CFTC) централизованная prediction market.

### Ключевые особенности
- **Settlement**: Централизованная, USD
- **Торговля**: Order book
- **Валюта**: USD (банковский перевод)
- **API**: REST
- **Ограничения**: KYC обязателен, доступ только из US

### API
- Docs: [trading-api.readme.kalshi.com](https://trading-api.readme.kalshi.com/)
- Требует отдельную API-key авторизацию

---

## Predict.fun

DEX prediction market на Blast L2.

### Ключевые особенности
- **Settlement**: On-chain (Blast)
- **Торговля**: AMM + Order book
- **Нет публичного API/SDK** для программной торговли

---

## Opinion, Probable, Limitless

Меньшие DEX prediction markets без публичных API.

- Данные доступны через Premarket.me агрегатор
- Программная торговля пока невозможна
- Арбитраж с этими платформами — только ручной (бот покажет ссылки)

---

## Матрица арбитражных пар

Бот находит арбитражные пары между любыми двумя платформами:

```
              polymarket  kalshi  predictfun  opinion  probable  limitless
polymarket        —         ✅       ✅         ✅       ✅        ✅
kalshi           ✅          —       ✅         ✅       ✅        ✅
predictfun       ✅         ✅        —         ✅       ✅        ✅
opinion          ✅         ✅       ✅          —       ✅        ✅
probable         ✅         ✅       ✅         ✅        —        ✅
limitless        ✅         ✅       ✅         ✅       ✅         —
```

**Автоматическое исполнение** возможно только для пар, где хотя бы одна сторона — Polymarket (через CLOB API).
