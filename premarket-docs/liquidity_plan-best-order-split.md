Title: Plan best order split - Premarket Docs

Description: Calculates how to split one buy or sell order across platforms to get a better execution. Send token ids for the same outcome across platforms and either amount_usd or amount_tokens. The response shows recommended allocation per platform and expected average execution price.

Source: https://docs.premarket.me/api-reference/liquidity/plan-best-order-split

---

[Skip to main content](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#content-area)
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
- liquidity[POSTPlan best order split](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split)
- [POSTPlan best order split](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split)
- market
- orderbook
- positions
- restricts
- arbitrage
- user
- v2/platforms
- v2/market
- [POSTPlan best order split](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split)
[POSTPlan best order split](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split)
cURL

```
curl --request POST \ --url https://api.example.com/api/liquidity/distribution \ --header 'Content-Type: application/json' \ --data ' { "tokens": {}, "direction": "<string>", "amount_usd": 100, "amount_tokens": 250, "include_fees": true, "enforce_min_orders": true, "optimize_gas": false, "gas_cost_per_source": 5 } '
```


```
curl --request POST \ --url https://api.example.com/api/liquidity/distribution \ --header 'Content-Type: application/json' \ --data ' { "tokens": {}, "direction": "<string>", "amount_usd": 100, "amount_tokens": 250, "include_fees": true, "enforce_min_orders": true, "optimize_gas": false, "gas_cost_per_source": 5 } '
```

200
success_buy

```
{ "combined": { "platform": "combined", "amount": 100, "price": 0.5271 }, "platforms": { "opinion": { "platform": "opinion", "used": true, "amount": 60, "price": 0.52, "best_price": 0.52, "hypothetical_tokens": 115.38, "hypothetical_price": 0.52, "available_liquidity": 1900, "liquidity_exhausted": false, "message_price": "52.0", "message_amount": "115.38" }, "predictfun": { "platform": "predictfun", "used": true, "amount": 40, "price": 0.535, "best_price": 0.531, "hypothetical_tokens": 74.77, "hypothetical_price": 0.535, "available_liquidity": 1200, "liquidity_exhausted": false, "message_price": "53.5", "message_amount": "74.77" } }, "distributions": [ { "platform": "opinion", "amount": 60, "price": 0.52 }, { "platform": "predictfun", "amount": 40, "price": 0.535 } ], "avg_price": 0.5271 }
```


```
{ "combined": { "platform": "combined", "amount": 100, "price": 0.5271 }, "platforms": { "opinion": { "platform": "opinion", "used": true, "amount": 60, "price": 0.52, "best_price": 0.52, "hypothetical_tokens": 115.38, "hypothetical_price": 0.52, "available_liquidity": 1900, "liquidity_exhausted": false, "message_price": "52.0", "message_amount": "115.38" }, "predictfun": { "platform": "predictfun", "used": true, "amount": 40, "price": 0.535, "best_price": 0.531, "hypothetical_tokens": 74.77, "hypothetical_price": 0.535, "available_liquidity": 1200, "liquidity_exhausted": false, "message_price": "53.5", "message_amount": "74.77" } }, "distributions": [ { "platform": "opinion", "amount": 60, "price": 0.52 }, { "platform": "predictfun", "amount": 40, "price": 0.535 } ], "avg_price": 0.5271 }
```

Calculates how to split one buy or sell order across platforms to get a better execution.
Send token ids for the same outcome across platforms and either amount_usd or amount_tokens. The response shows recommended allocation per platform and expected average execution price.

```
amount_usd
```


```
amount_tokens
```

cURL

```
curl --request POST \ --url https://api.example.com/api/liquidity/distribution \ --header 'Content-Type: application/json' \ --data ' { "tokens": {}, "direction": "<string>", "amount_usd": 100, "amount_tokens": 250, "include_fees": true, "enforce_min_orders": true, "optimize_gas": false, "gas_cost_per_source": 5 } '
```


```
curl --request POST \ --url https://api.example.com/api/liquidity/distribution \ --header 'Content-Type: application/json' \ --data ' { "tokens": {}, "direction": "<string>", "amount_usd": 100, "amount_tokens": 250, "include_fees": true, "enforce_min_orders": true, "optimize_gas": false, "gas_cost_per_source": 5 } '
```

200
success_buy

```
{ "combined": { "platform": "combined", "amount": 100, "price": 0.5271 }, "platforms": { "opinion": { "platform": "opinion", "used": true, "amount": 60, "price": 0.52, "best_price": 0.52, "hypothetical_tokens": 115.38, "hypothetical_price": 0.52, "available_liquidity": 1900, "liquidity_exhausted": false, "message_price": "52.0", "message_amount": "115.38" }, "predictfun": { "platform": "predictfun", "used": true, "amount": 40, "price": 0.535, "best_price": 0.531, "hypothetical_tokens": 74.77, "hypothetical_price": 0.535, "available_liquidity": 1200, "liquidity_exhausted": false, "message_price": "53.5", "message_amount": "74.77" } }, "distributions": [ { "platform": "opinion", "amount": 60, "price": 0.52 }, { "platform": "predictfun", "amount": 40, "price": 0.535 } ], "avg_price": 0.5271 }
```


```
{ "combined": { "platform": "combined", "amount": 100, "price": 0.5271 }, "platforms": { "opinion": { "platform": "opinion", "used": true, "amount": 60, "price": 0.52, "best_price": 0.52, "hypothetical_tokens": 115.38, "hypothetical_price": 0.52, "available_liquidity": 1900, "liquidity_exhausted": false, "message_price": "52.0", "message_amount": "115.38" }, "predictfun": { "platform": "predictfun", "used": true, "amount": 40, "price": 0.535, "best_price": 0.531, "hypothetical_tokens": 74.77, "hypothetical_price": 0.535, "available_liquidity": 1200, "liquidity_exhausted": false, "message_price": "53.5", "message_amount": "74.77" } }, "distributions": [ { "platform": "opinion", "amount": 60, "price": 0.52 }, { "platform": "predictfun", "amount": 40, "price": 0.535 } ], "avg_price": 0.5271 }
```

Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Body
Request for liquidity distribution
[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#body-tokens)
Token identifiers for the same outcome across platforms. Supported keys: polymarket, opinion, probable, predictfun, limitless. Recommended format is object with token_id (and market_id where applicable).

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


```
token_id
```


```
market_id
```

Show child attributes

```
{ "opinion": { "outcome": "yes", "token_id": "987654" }, "polymarket": { "outcome": "yes", "token_id": "123456" }, "predictfun": { "market_id": "991122", "outcome": "yes", "token_id": "555" }}
```


```
{ "opinion": { "outcome": "yes", "token_id": "987654" }, "polymarket": { "outcome": "yes", "token_id": "123456" }, "predictfun": { "market_id": "991122", "outcome": "yes", "token_id": "555" }}
```

[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#body-direction)
Trade direction: buy or sell.

```
buy
```


```
sell
```

"buy"

```
"buy"
```

[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#body-amount-usd-one-of-0)
Order size in USD. Provide either amount_usd or amount_tokens.

```
amount_usd
```


```
amount_tokens
```


```
x > 0
```

100

```
100
```

[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#body-amount-tokens-one-of-0)
Order size in tokens. Provide either amount_usd or amount_tokens.

```
amount_usd
```


```
amount_tokens
```


```
x > 0
```

250

```
250
```

[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#body-include-fees)
Include platform fees in price calculation.
[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#body-enforce-min-orders)
Apply platform minimum order constraints.
[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#body-optimize-gas)
Prefer fewer sources when estimating execution.
[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#body-gas-cost-per-source)
Estimated gas cost per additional source in USD.

#### Response
Optimal distribution calculated
Response with optimal liquidity distribution
[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#response-combined)
Distribution for a single platform
Show child attributes
[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#response-platforms)
Show child attributes
[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#response-distributions)
Show child attributes
[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#response-avg-price)
[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#response-orderbook-one-of-0)
[​](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split#response-error-one-of-0)
[Get market token balances](https://docs.premarket.me/api-reference/balances/get-market-token-balances)
[Get market by id](https://docs.premarket.me/api-reference/market/get-market-by-id)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

