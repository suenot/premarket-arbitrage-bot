Title: Get market token balances - Premarket Docs

Description: Returns non-zero balances for market tokens held by the provided wallets. Use token_ids when you want balances for specific outcomes only. If token_ids is omitted, the endpoint returns balances based on wallet activity.

Source: https://docs.premarket.me/api-reference/balances/get-market-token-balances

---

[Skip to main content](https://docs.premarket.me/api-reference/balances/get-market-token-balances#content-area)
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
curl --request GET \ --url https://api.example.com/api/balances/erc1155
```


```
curl --request GET \ --url https://api.example.com/api/balances/erc1155
```


```
{ "balances": [ { "platform": "<string>", "contract_address": "<string>", "token_id": "<string>", "balance": 123 } ], "addresses": [ "<string>" ], "error": "<string>" }
```


```
{ "balances": [ { "platform": "<string>", "contract_address": "<string>", "token_id": "<string>", "balance": 123 } ], "addresses": [ "<string>" ], "error": "<string>" }
```


# Get market token balances
Returns non-zero balances for market tokens held by the provided wallets.
Use token_ids when you want balances for specific outcomes only. If token_ids is omitted, the endpoint returns balances based on wallet activity.

```
token_ids
```


```
token_ids
```

cURL

```
curl --request GET \ --url https://api.example.com/api/balances/erc1155
```


```
curl --request GET \ --url https://api.example.com/api/balances/erc1155
```


```
{ "balances": [ { "platform": "<string>", "contract_address": "<string>", "token_id": "<string>", "balance": 123 } ], "addresses": [ "<string>" ], "error": "<string>" }
```


```
{ "balances": [ { "platform": "<string>", "contract_address": "<string>", "token_id": "<string>", "balance": 123 } ], "addresses": [ "<string>" ], "error": "<string>" }
```


## Documentation Index
Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Query Parameters
[​](https://docs.premarket.me/api-reference/balances/get-market-token-balances#parameter-addresses)
Comma-separated wallet addresses to inspect.
[​](https://docs.premarket.me/api-reference/balances/get-market-token-balances#parameter-token-ids)
Optional comma-separated token ids to limit the result.

#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/balances/get-market-token-balances#response-balances)
ERC1155 balances for requested addresses/tokens
Show child attributes
[​](https://docs.premarket.me/api-reference/balances/get-market-token-balances#response-addresses)
Input wallet addresses
[​](https://docs.premarket.me/api-reference/balances/get-market-token-balances#response-error-one-of-0)
Error message when request cannot be processed
[Get wallet cash balances](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances)
[Plan best order split](https://docs.premarket.me/api-reference/liquidity/plan-best-order-split)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

