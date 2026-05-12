Title: Get outcome orderbook - Premarket Docs

Description: Returns price levels for one outcome across available platforms, plus a combined book. You can pass token ids for YES/NO sides (or market id for supported platforms) and use the response to show depth, best price, and cross-platform comparison.

Source: https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook

---

[Skip to main content](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#content-area)
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
- orderbook[GETGet outcome orderbook](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook)
- [GETGet outcome orderbook](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook)
- positions
- restricts
- arbitrage
- user
- v2/platforms
- v2/market
- [GETGet outcome orderbook](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook)
[GETGet outcome orderbook](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook)
cURL

```
curl --request GET \ --url https://api.example.com/api/orderbook
```


```
curl --request GET \ --url https://api.example.com/api/orderbook
```


```
{ "polymarket": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "opinion": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "probable": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "predictfun": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "limitless": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "combined": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 } }
```


```
{ "polymarket": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "opinion": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "probable": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "predictfun": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "limitless": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "combined": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 } }
```

Returns price levels for one outcome across available platforms, plus a combined book.
You can pass token ids for YES/NO sides (or market id for supported platforms) and use the response to show depth, best price, and cross-platform comparison.
cURL

```
curl --request GET \ --url https://api.example.com/api/orderbook
```


```
curl --request GET \ --url https://api.example.com/api/orderbook
```


```
{ "polymarket": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "opinion": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "probable": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "predictfun": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "limitless": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "combined": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 } }
```


```
{ "polymarket": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "opinion": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "probable": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "predictfun": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "limitless": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 }, "combined": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "no": { "asks": [ [ 123 ] ], "bids": [ [ 123 ] ] }, "yes_price": 123, "no_price": 123 } }
```

Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Query Parameters
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#parameter-pm-yes-token)
Polymarket YES token id.
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#parameter-pm-no-token)
Polymarket NO token id.
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#parameter-op-yes-token)
Opinion YES token id.
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#parameter-op-no-token)
Opinion NO token id.
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#parameter-prob-yes-token)
Probable YES token id.
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#parameter-prob-no-token)
Probable NO token id.
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#parameter-pf-market-id)
Predictfun market id.
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#parameter-lim-market-slug)
Limitless market slug.
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#parameter-lim-yes-token)
Limitless YES token id.
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#parameter-lim-no-token)
Limitless NO token id.

#### Response
Successful Response
Combined orderbook response
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#response-polymarket)
Normalized platform-specific orderbook
Show child attributes
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#response-opinion)
Normalized platform-specific orderbook
Show child attributes
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#response-probable)
Normalized platform-specific orderbook
Show child attributes
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#response-predictfun)
Normalized platform-specific orderbook
Show child attributes
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#response-limitless)
Normalized platform-specific orderbook
Show child attributes
[​](https://docs.premarket.me/api-reference/orderbook/get-outcome-orderbook#response-combined)
Normalized platform-specific orderbook
Show child attributes
[Get market by id](https://docs.premarket.me/api-reference/market/get-market-by-id)
[Get portfolio positions](https://docs.premarket.me/api-reference/positions/get-portfolio-positions)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

