Title: Get market by id - Premarket Docs

Description: Returns one market with its outcomes and event context for the selected platform. How to get market_id: - Use GET /api/event/{slug} and take the platform-specific market id from outcomes (for example polymarket_market_id, opinion_market_id, etc.). Auth: - Requires X-API-Key header.

Source: https://docs.premarket.me/api-reference/v2market/get-market-by-id

---

[Skip to main content](https://docs.premarket.me/api-reference/v2market/get-market-by-id#content-area)
[Premarket Docs home page](https://docs.premarket.me/)
- [X (Twitter)](https://x.com/pre_markets)
[X (Twitter)](https://x.com/pre_markets)

##### Overview
- [Premarket Docs](https://docs.premarket.me/)
- [Getting Started](https://docs.premarket.me/getting-started)
[Premarket Docs](https://docs.premarket.me/)
[Getting Started](https://docs.premarket.me/getting-started)

##### API Reference
- events
- auth
- balances
- liquidity
- market
- orderbook
- positions
- restricts
- arbitrage
- user
- v2/platforms
- v2/market[GETGet market by id](https://docs.premarket.me/api-reference/v2market/get-market-by-id)
- [GETGet market by id](https://docs.premarket.me/api-reference/v2market/get-market-by-id)
- [GETGet market by id](https://docs.premarket.me/api-reference/v2market/get-market-by-id)
[GETGet market by id](https://docs.premarket.me/api-reference/v2market/get-market-by-id)
cURL

```
curl --request GET \ --url https://api.example.com/api/v2/market/id
```


```
curl --request GET \ --url https://api.example.com/api/v2/market/id
```


```
{ "market": { "market_id": "<string>", "source": "polymarket", "market_slug": "<string>", "question": "<string>", "description": "<string>", "cache": false, "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123, "orderbook": { "spread": 0, "lastTradePrice": 0, "bestBid": 0, "bestAsk": 0 } } ] }, "event": { "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "event_icon": "<string>", "end_date": "<string>" }, "series": { "ticker": "<string>", "slug": "<string>" }, "related_markets": [ { "market_id": "<string>", "source": "polymarket", "event": { "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "event_icon": "<string>", "end_date": "<string>" }, "market_slug": "<string>", "question": "<string>", "description": "<string>", "cache": false, "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123, "orderbook": { "spread": 0, "lastTradePrice": 0, "bestBid": 0, "bestAsk": 0 } } ], "series": { "ticker": "<string>", "slug": "<string>" } } ] }
```


```
{ "market": { "market_id": "<string>", "source": "polymarket", "market_slug": "<string>", "question": "<string>", "description": "<string>", "cache": false, "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123, "orderbook": { "spread": 0, "lastTradePrice": 0, "bestBid": 0, "bestAsk": 0 } } ] }, "event": { "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "event_icon": "<string>", "end_date": "<string>" }, "series": { "ticker": "<string>", "slug": "<string>" }, "related_markets": [ { "market_id": "<string>", "source": "polymarket", "event": { "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "event_icon": "<string>", "end_date": "<string>" }, "market_slug": "<string>", "question": "<string>", "description": "<string>", "cache": false, "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123, "orderbook": { "spread": 0, "lastTradePrice": 0, "bestBid": 0, "bestAsk": 0 } } ], "series": { "ticker": "<string>", "slug": "<string>" } } ] }
```

Returns one market with its outcomes and event context for the selected platform.
How to get market_id:

```
market_id
```

- Use GET /api/event/{slug} and take the platform-specific market id from outcomes
(for example polymarket_market_id, opinion_market_id, etc.).

```
GET /api/event/{slug}
```


```
polymarket_market_id
```


```
opinion_market_id
```

Auth:
- Requires X-API-Key header.

```
X-API-Key
```

cURL

```
curl --request GET \ --url https://api.example.com/api/v2/market/id
```


```
curl --request GET \ --url https://api.example.com/api/v2/market/id
```


```
{ "market": { "market_id": "<string>", "source": "polymarket", "market_slug": "<string>", "question": "<string>", "description": "<string>", "cache": false, "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123, "orderbook": { "spread": 0, "lastTradePrice": 0, "bestBid": 0, "bestAsk": 0 } } ] }, "event": { "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "event_icon": "<string>", "end_date": "<string>" }, "series": { "ticker": "<string>", "slug": "<string>" }, "related_markets": [ { "market_id": "<string>", "source": "polymarket", "event": { "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "event_icon": "<string>", "end_date": "<string>" }, "market_slug": "<string>", "question": "<string>", "description": "<string>", "cache": false, "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123, "orderbook": { "spread": 0, "lastTradePrice": 0, "bestBid": 0, "bestAsk": 0 } } ], "series": { "ticker": "<string>", "slug": "<string>" } } ] }
```


```
{ "market": { "market_id": "<string>", "source": "polymarket", "market_slug": "<string>", "question": "<string>", "description": "<string>", "cache": false, "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123, "orderbook": { "spread": 0, "lastTradePrice": 0, "bestBid": 0, "bestAsk": 0 } } ] }, "event": { "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "event_icon": "<string>", "end_date": "<string>" }, "series": { "ticker": "<string>", "slug": "<string>" }, "related_markets": [ { "market_id": "<string>", "source": "polymarket", "event": { "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "event_icon": "<string>", "end_date": "<string>" }, "market_slug": "<string>", "question": "<string>", "description": "<string>", "cache": false, "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123, "orderbook": { "spread": 0, "lastTradePrice": 0, "bestBid": 0, "bestAsk": 0 } } ], "series": { "ticker": "<string>", "slug": "<string>" } } ] }
```

Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Query Parameters
[​](https://docs.premarket.me/api-reference/v2market/get-market-by-id#parameter-source)

```
polymarket
```


```
opinion
```


```
probable
```


```
predictfun
```


```
limitless
```

[​](https://docs.premarket.me/api-reference/v2market/get-market-by-id#parameter-market-id)
[​](https://docs.premarket.me/api-reference/v2market/get-market-by-id#parameter-cache)

#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/v2market/get-market-by-id#response-market)
Show child attributes
[​](https://docs.premarket.me/api-reference/v2market/get-market-by-id#response-event)
Show child attributes
[​](https://docs.premarket.me/api-reference/v2market/get-market-by-id#response-series-one-of-0)
Show child attributes
[​](https://docs.premarket.me/api-reference/v2market/get-market-by-id#response-related-markets)
Show child attributes
[List platforms](https://docs.premarket.me/api-reference/v2platforms/list-platforms)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

