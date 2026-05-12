Title: Get market by id - Premarket Docs

Description: Returns one market with its outcomes and event context for the selected platform. How to get market_id: - Use GET /api/event/{slug} and take the platform-specific market id from outcomes (for example polymarket_market_id, opinion_market_id, etc.). Auth: - Requires X-API-Key header.

Source: https://docs.premarket.me/api-reference/market/get-market-by-id

---

[Skip to main content](https://docs.premarket.me/api-reference/market/get-market-by-id#content-area)
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
- market[GETGet market by id](https://docs.premarket.me/api-reference/market/get-market-by-id)
- [GETGet market by id](https://docs.premarket.me/api-reference/market/get-market-by-id)
- orderbook
- positions
- restricts
- arbitrage
- user
- v2/platforms
- v2/market
- [GETGet market by id](https://docs.premarket.me/api-reference/market/get-market-by-id)
[GETGet market by id](https://docs.premarket.me/api-reference/market/get-market-by-id)
cURL

```
curl --request GET \ --url https://api.example.com/api/market/{market_id} \ --header 'X-API-Key: <api-key>'
```


```
curl --request GET \ --url https://api.example.com/api/market/{market_id} \ --header 'X-API-Key: <api-key>'
```


```
{ "market_id": "<string>", "source": "<string>", "market_slug": "<string>", "question": "<string>", "description": "<string>", "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "end_date": "<string>", "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123 } ] }
```


```
{ "market_id": "<string>", "source": "<string>", "market_slug": "<string>", "question": "<string>", "description": "<string>", "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "end_date": "<string>", "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123 } ] }
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
curl --request GET \ --url https://api.example.com/api/market/{market_id} \ --header 'X-API-Key: <api-key>'
```


```
curl --request GET \ --url https://api.example.com/api/market/{market_id} \ --header 'X-API-Key: <api-key>'
```


```
{ "market_id": "<string>", "source": "<string>", "market_slug": "<string>", "question": "<string>", "description": "<string>", "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "end_date": "<string>", "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123 } ] }
```


```
{ "market_id": "<string>", "source": "<string>", "market_slug": "<string>", "question": "<string>", "description": "<string>", "event_id": "<string>", "event_slug": "<string>", "event_title": "<string>", "end_date": "<string>", "options": [ { "outcome": "<string>", "token_id": "<string>", "price": 123 } ] }
```

Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Authorizations
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#authorization-x-api-key)

#### Path Parameters
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#parameter-market-id)
Market id from event details.

#### Query Parameters
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#parameter-source)

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


#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#response-market-id)
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#response-source)
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#response-market-slug-one-of-0)
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#response-question-one-of-0)
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#response-description-one-of-0)
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#response-event-id-one-of-0)
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#response-event-slug-one-of-0)
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#response-event-title-one-of-0)
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#response-end-date-one-of-0)
[​](https://docs.premarket.me/api-reference/market/get-market-by-id#response-options)
Show child attributes
[Plan best order split](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split)
[Get outcome orderbook](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

