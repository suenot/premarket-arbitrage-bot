Title: Get wallet cash balances - Premarket Docs

Description: Returns stablecoin balances for one or more wallet addresses. Use this endpoint to show available cash before trading or to build a portfolio header view. Response includes per-chain totals (polygon, bsc, base), overall total, and per-address balances inside each chain object.

Source: https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances

---

[Skip to main content](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances#content-area)
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
- balances[GETGet wallet cash balances](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances)[GETGet market token balances](https://docs.premarket.me/api-reference/balances/get-market-token-balances)
- [GETGet wallet cash balances](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances)
- [GETGet market token balances](https://docs.premarket.me/api-reference/balances/get-market-token-balances)
- liquidity
- market
- orderbook
- positions
- restricts
- arbitrage
- user
- v2/platforms
- v2/market
- [GETGet wallet cash balances](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances)
- [GETGet market token balances](https://docs.premarket.me/api-reference/balances/get-market-token-balances)
[GETGet wallet cash balances](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances)
[GETGet market token balances](https://docs.premarket.me/api-reference/balances/get-market-token-balances)
cURL

```
curl --request GET \ --url https://api.example.com/api/balances
```


```
curl --request GET \ --url https://api.example.com/api/balances
```


```
{ "balances": { "polygon": { "balance": 0, "symbol": "", "addresses": {} }, "bsc": { "balance": 0, "symbol": "", "addresses": {} }, "base": { "balance": 0, "symbol": "", "addresses": {} }, "total": 0 }, "addresses": [ "<string>" ], "error": "<string>" }
```


```
{ "balances": { "polygon": { "balance": 0, "symbol": "", "addresses": {} }, "bsc": { "balance": 0, "symbol": "", "addresses": {} }, "base": { "balance": 0, "symbol": "", "addresses": {} }, "total": 0 }, "addresses": [ "<string>" ], "error": "<string>" }
```

Returns stablecoin balances for one or more wallet addresses.
Use this endpoint to show available cash before trading or to build a portfolio header view.
Response includes per-chain totals (polygon, bsc, base), overall total, and per-address balances inside each chain object.

```
polygon
```


```
bsc
```


```
base
```


```
total
```

cURL

```
curl --request GET \ --url https://api.example.com/api/balances
```


```
curl --request GET \ --url https://api.example.com/api/balances
```


```
{ "balances": { "polygon": { "balance": 0, "symbol": "", "addresses": {} }, "bsc": { "balance": 0, "symbol": "", "addresses": {} }, "base": { "balance": 0, "symbol": "", "addresses": {} }, "total": 0 }, "addresses": [ "<string>" ], "error": "<string>" }
```


```
{ "balances": { "polygon": { "balance": 0, "symbol": "", "addresses": {} }, "bsc": { "balance": 0, "symbol": "", "addresses": {} }, "base": { "balance": 0, "symbol": "", "addresses": {} }, "total": 0 }, "addresses": [ "<string>" ], "error": "<string>" }
```

Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Query Parameters
[​](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances#parameter-addresses)
Comma-separated wallet addresses, for example: 0xabc...,0xdef....

```
0xabc...,0xdef...
```


#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances#response-balances)
Aggregated stablecoin balances
Show child attributes
[​](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances#response-addresses)
Input wallet addresses from the request.

```
["0x1111...aaaa", "0x2222...bbbb"]
```


```
["0x1111...aaaa", "0x2222...bbbb"]
```

[​](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances#response-error-one-of-0)
Error message when request cannot be processed
[Create API key](https://docs.premarket.me/api-reference/auth/create-api-key)
[Get market token balances](https://docs.premarket.me/api-reference/balances/get-market-token-balances)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

