Title: Get portfolio positions - Premarket Docs

Description: Returns open positions for one or more wallets in a single portfolio view. Includes position list, aggregated totals, and related wallet addresses for supported platforms. Use connected EOA wallet addresses in addresses (comma-separated).

Source: https://docs.premarket.me/api-reference/positions/get-portfolio-positions

---

[Skip to main content](https://docs.premarket.me/api-reference/positions/get-portfolio-positions#content-area)
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
- positions[GETGet portfolio positions](https://docs.premarket.me/api-reference/positions/get-portfolio-positions)
- [GETGet portfolio positions](https://docs.premarket.me/api-reference/positions/get-portfolio-positions)
- restricts
- arbitrage
- user
- v2/platforms
- v2/market
- [GETGet portfolio positions](https://docs.premarket.me/api-reference/positions/get-portfolio-positions)
[GETGet portfolio positions](https://docs.premarket.me/api-reference/positions/get-portfolio-positions)
cURL

```
curl --request GET \ --url https://api.example.com/api/positions/aggregated
```


```
curl --request GET \ --url https://api.example.com/api/positions/aggregated
```


```
{ "summary": { "total_value": 0, "total_positions": 0, "active_positions": 0, "claimable_value": 0, "by_platform": { "polymarket": 0, "opinion": 0, "probable": 0, "predictfun": 0, "limitless": 0 }, "by_chain": { "polygon": 0, "bsc": 0, "base": 0 } }, "balances": { "total": 0, "polygon": { "balance": 0, "symbol": "", "addresses": {} }, "bsc": { "balance": 0, "symbol": "", "addresses": {} }, "base": { "balance": 0, "symbol": "", "addresses": {} } }, "positions": [ { "token_id": "<string>", "event_title": "<string>", "source": "<string>", "outcome": "<string>", "market_id": 123, "chain_id": "<string>", "contract_address": "<string>", "question": "<string>", "event_id": "<string>", "event_slug": "<string>", "condition_id": "<string>", "balance_raw": 0, "balance": 0, "current_price": 0, "value": 0, "is_active": false, "market_status": "<string>", "claim_available": false, "claim_params": { "collateral_token": "<string>", "parent_collection_id": "<string>", "condition_id": "<string>", "index_sets": [ 123 ] } } ], "addresses": [ "<string>" ], "wallets": [ { "eoa": "<string>", "wallets": { "polymarket": { "safe_address": "<string>", "error": "<string>" }, "opinion": { "safe_address": "<string>", "error": "<string>" } } } ] }
```


```
{ "summary": { "total_value": 0, "total_positions": 0, "active_positions": 0, "claimable_value": 0, "by_platform": { "polymarket": 0, "opinion": 0, "probable": 0, "predictfun": 0, "limitless": 0 }, "by_chain": { "polygon": 0, "bsc": 0, "base": 0 } }, "balances": { "total": 0, "polygon": { "balance": 0, "symbol": "", "addresses": {} }, "bsc": { "balance": 0, "symbol": "", "addresses": {} }, "base": { "balance": 0, "symbol": "", "addresses": {} } }, "positions": [ { "token_id": "<string>", "event_title": "<string>", "source": "<string>", "outcome": "<string>", "market_id": 123, "chain_id": "<string>", "contract_address": "<string>", "question": "<string>", "event_id": "<string>", "event_slug": "<string>", "condition_id": "<string>", "balance_raw": 0, "balance": 0, "current_price": 0, "value": 0, "is_active": false, "market_status": "<string>", "claim_available": false, "claim_params": { "collateral_token": "<string>", "parent_collection_id": "<string>", "condition_id": "<string>", "index_sets": [ 123 ] } } ], "addresses": [ "<string>" ], "wallets": [ { "eoa": "<string>", "wallets": { "polymarket": { "safe_address": "<string>", "error": "<string>" }, "opinion": { "safe_address": "<string>", "error": "<string>" } } } ] }
```

Returns open positions for one or more wallets in a single portfolio view.
Includes position list, aggregated totals, and related wallet addresses for supported platforms.
Use connected EOA wallet addresses in addresses (comma-separated).

```
addresses
```

cURL

```
curl --request GET \ --url https://api.example.com/api/positions/aggregated
```


```
curl --request GET \ --url https://api.example.com/api/positions/aggregated
```


```
{ "summary": { "total_value": 0, "total_positions": 0, "active_positions": 0, "claimable_value": 0, "by_platform": { "polymarket": 0, "opinion": 0, "probable": 0, "predictfun": 0, "limitless": 0 }, "by_chain": { "polygon": 0, "bsc": 0, "base": 0 } }, "balances": { "total": 0, "polygon": { "balance": 0, "symbol": "", "addresses": {} }, "bsc": { "balance": 0, "symbol": "", "addresses": {} }, "base": { "balance": 0, "symbol": "", "addresses": {} } }, "positions": [ { "token_id": "<string>", "event_title": "<string>", "source": "<string>", "outcome": "<string>", "market_id": 123, "chain_id": "<string>", "contract_address": "<string>", "question": "<string>", "event_id": "<string>", "event_slug": "<string>", "condition_id": "<string>", "balance_raw": 0, "balance": 0, "current_price": 0, "value": 0, "is_active": false, "market_status": "<string>", "claim_available": false, "claim_params": { "collateral_token": "<string>", "parent_collection_id": "<string>", "condition_id": "<string>", "index_sets": [ 123 ] } } ], "addresses": [ "<string>" ], "wallets": [ { "eoa": "<string>", "wallets": { "polymarket": { "safe_address": "<string>", "error": "<string>" }, "opinion": { "safe_address": "<string>", "error": "<string>" } } } ] }
```


```
{ "summary": { "total_value": 0, "total_positions": 0, "active_positions": 0, "claimable_value": 0, "by_platform": { "polymarket": 0, "opinion": 0, "probable": 0, "predictfun": 0, "limitless": 0 }, "by_chain": { "polygon": 0, "bsc": 0, "base": 0 } }, "balances": { "total": 0, "polygon": { "balance": 0, "symbol": "", "addresses": {} }, "bsc": { "balance": 0, "symbol": "", "addresses": {} }, "base": { "balance": 0, "symbol": "", "addresses": {} } }, "positions": [ { "token_id": "<string>", "event_title": "<string>", "source": "<string>", "outcome": "<string>", "market_id": 123, "chain_id": "<string>", "contract_address": "<string>", "question": "<string>", "event_id": "<string>", "event_slug": "<string>", "condition_id": "<string>", "balance_raw": 0, "balance": 0, "current_price": 0, "value": 0, "is_active": false, "market_status": "<string>", "claim_available": false, "claim_params": { "collateral_token": "<string>", "parent_collection_id": "<string>", "condition_id": "<string>", "index_sets": [ 123 ] } } ], "addresses": [ "<string>" ], "wallets": [ { "eoa": "<string>", "wallets": { "polymarket": { "safe_address": "<string>", "error": "<string>" }, "opinion": { "safe_address": "<string>", "error": "<string>" } } } ] }
```

Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Query Parameters
[​](https://docs.premarket.me/api-reference/positions/get-portfolio-positions#parameter-addresses)
Comma-separated wallet addresses to include in portfolio view.

#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/positions/get-portfolio-positions#response-summary)
Aggregated summary
Show child attributes

```
{ "active_positions": 13, "by_chain": { "base": 40, "bsc": 367.29, "polygon": 841.12 }, "by_platform": { "limitless": 40, "opinion": 312.44, "polymarket": 801.12, "predictfun": 0, "probable": 94.85 }, "claimable_value": 57.22, "total_positions": 18, "total_value": 1248.41}
```


```
{ "active_positions": 13, "by_chain": { "base": 40, "bsc": 367.29, "polygon": 841.12 }, "by_platform": { "limitless": 40, "opinion": 312.44, "polymarket": 801.12, "predictfun": 0, "probable": 94.85 }, "claimable_value": 57.22, "total_positions": 18, "total_value": 1248.41}
```

[​](https://docs.premarket.me/api-reference/positions/get-portfolio-positions#response-balances)
Stablecoin balances for all queried wallets
Show child attributes
[​](https://docs.premarket.me/api-reference/positions/get-portfolio-positions#response-positions)
Position list
Show child attributes
[​](https://docs.premarket.me/api-reference/positions/get-portfolio-positions#response-addresses)
Queried EOA addresses
[​](https://docs.premarket.me/api-reference/positions/get-portfolio-positions#response-wallets)
Safe wallets per EOA
Show child attributes
[Get outcome orderbook](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook)
[Check regional availability](https://docs.premarket.me/api-reference/restricts/check-regional-availability)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

