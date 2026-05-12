Title: Get arbitrage opportunities - Premarket Docs

Description: Returns current arbitrage opportunities between platforms. Each item includes expected spread, estimated profit, available depth, and market links. Use limit and offset for pagination.

Source: https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities

---

[Skip to main content](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities#content-area)
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
- arbitrage[GETGet arbitrage opportunities](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities)
- [GETGet arbitrage opportunities](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities)
- user
- v2/platforms
- v2/market
- [GETGet arbitrage opportunities](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities)
[GETGet arbitrage opportunities](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities)
cURL

```
curl --request GET \ --url https://api.example.com/api/arbitrage \ --header 'Authorization: Bearer <token>'
```


```
curl --request GET \ --url https://api.example.com/api/arbitrage \ --header 'Authorization: Bearer <token>'
```


```
{ "pairs": [ { "platform1": { "source": "<string>", "market_id": 123, "event_id": "<string>", "event_title": "<string>", "question": "<string>", "description": "<string>", "url": "<string>", "yes_price": 123, "no_price": 123, "yes_token_id": "<string>", "no_token_id": "<string>" }, "platform2": { "source": "<string>", "market_id": 123, "event_id": "<string>", "event_title": "<string>", "question": "<string>", "description": "<string>", "url": "<string>", "yes_price": 123, "no_price": 123, "yes_token_id": "<string>", "no_token_id": "<string>" }, "arbitrage": { "buy_yes_price": 0, "buy_yes_source": "", "buy_no_price": 0, "buy_no_source": "", "total_cost": 0, "profit": 0, "profit_pct": 0, "days_until": 0, "apr": 0, "platform_fee": 123 }, "depth": { "max_shares": 0, "max_usd": 0, "yes_usd": 0, "no_usd": 0, "avg_yes": 0, "avg_no": 0 }, "end_date": "<string>" } ], "total": 0, "count": 0 }
```


```
{ "pairs": [ { "platform1": { "source": "<string>", "market_id": 123, "event_id": "<string>", "event_title": "<string>", "question": "<string>", "description": "<string>", "url": "<string>", "yes_price": 123, "no_price": 123, "yes_token_id": "<string>", "no_token_id": "<string>" }, "platform2": { "source": "<string>", "market_id": 123, "event_id": "<string>", "event_title": "<string>", "question": "<string>", "description": "<string>", "url": "<string>", "yes_price": 123, "no_price": 123, "yes_token_id": "<string>", "no_token_id": "<string>" }, "arbitrage": { "buy_yes_price": 0, "buy_yes_source": "", "buy_no_price": 0, "buy_no_source": "", "total_cost": 0, "profit": 0, "profit_pct": 0, "days_until": 0, "apr": 0, "platform_fee": 123 }, "depth": { "max_shares": 0, "max_usd": 0, "yes_usd": 0, "no_usd": 0, "avg_yes": 0, "avg_no": 0 }, "end_date": "<string>" } ], "total": 0, "count": 0 }
```

Returns current arbitrage opportunities between platforms.
Each item includes expected spread, estimated profit, available depth, and market links. Use limit and offset for pagination.

```
limit
```


```
offset
```

cURL

```
curl --request GET \ --url https://api.example.com/api/arbitrage \ --header 'Authorization: Bearer <token>'
```


```
curl --request GET \ --url https://api.example.com/api/arbitrage \ --header 'Authorization: Bearer <token>'
```


```
{ "pairs": [ { "platform1": { "source": "<string>", "market_id": 123, "event_id": "<string>", "event_title": "<string>", "question": "<string>", "description": "<string>", "url": "<string>", "yes_price": 123, "no_price": 123, "yes_token_id": "<string>", "no_token_id": "<string>" }, "platform2": { "source": "<string>", "market_id": 123, "event_id": "<string>", "event_title": "<string>", "question": "<string>", "description": "<string>", "url": "<string>", "yes_price": 123, "no_price": 123, "yes_token_id": "<string>", "no_token_id": "<string>" }, "arbitrage": { "buy_yes_price": 0, "buy_yes_source": "", "buy_no_price": 0, "buy_no_source": "", "total_cost": 0, "profit": 0, "profit_pct": 0, "days_until": 0, "apr": 0, "platform_fee": 123 }, "depth": { "max_shares": 0, "max_usd": 0, "yes_usd": 0, "no_usd": 0, "avg_yes": 0, "avg_no": 0 }, "end_date": "<string>" } ], "total": 0, "count": 0 }
```


```
{ "pairs": [ { "platform1": { "source": "<string>", "market_id": 123, "event_id": "<string>", "event_title": "<string>", "question": "<string>", "description": "<string>", "url": "<string>", "yes_price": 123, "no_price": 123, "yes_token_id": "<string>", "no_token_id": "<string>" }, "platform2": { "source": "<string>", "market_id": 123, "event_id": "<string>", "event_title": "<string>", "question": "<string>", "description": "<string>", "url": "<string>", "yes_price": 123, "no_price": 123, "yes_token_id": "<string>", "no_token_id": "<string>" }, "arbitrage": { "buy_yes_price": 0, "buy_yes_source": "", "buy_no_price": 0, "buy_no_source": "", "total_cost": 0, "profit": 0, "profit_pct": 0, "days_until": 0, "apr": 0, "platform_fee": 123 }, "depth": { "max_shares": 0, "max_usd": 0, "yes_usd": 0, "no_usd": 0, "avg_yes": 0, "avg_no": 0 }, "end_date": "<string>" } ], "total": 0, "count": 0 }
```

Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Authorizations
[​](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities#authorization-authorization)
Bearer authentication header of the form Bearer <token>, where <token> is your auth token.

```
Bearer <token>
```


```
<token>
```


#### Query Parameters
[​](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities#parameter-one-of-0)
Optional number of pairs to return.
[​](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities#parameter-offset)
Zero-based offset in the result list.

```
x >= 0
```

[​](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities#parameter-one-of-0)
Optional arbitrage id to force-include in guest preview.

#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities#response-pairs)
Arbitrage pairs list
Show child attributes
[​](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities#response-total)
Total candidate pairs before profitability filter
[​](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities#response-count)
Returned pairs count
[Check regional availability](https://docs.premarket.me/api-reference/restricts/check-regional-availability)
[Get my profile](https://docs.premarket.me/api-reference/user/get-my-profile)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

