Title: Check regional availability - Premarket Docs

Description: Returns whether trading is available for the current user region. Use this endpoint to decide if trading actions should be enabled, disabled, or limited per platform.

Source: https://docs.premarket.me/api-reference/restricts/check-regional-availability

---

[Skip to main content](https://docs.premarket.me/api-reference/restricts/check-regional-availability#content-area)
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
- restricts[GETCheck regional availability](https://docs.premarket.me/api-reference/restricts/check-regional-availability)
- [GETCheck regional availability](https://docs.premarket.me/api-reference/restricts/check-regional-availability)
- arbitrage
- user
- v2/platforms
- v2/market
- [GETCheck regional availability](https://docs.premarket.me/api-reference/restricts/check-regional-availability)
[GETCheck regional availability](https://docs.premarket.me/api-reference/restricts/check-regional-availability)
cURL

```
curl --request GET \ --url https://api.example.com/api/check_restricts
```


```
curl --request GET \ --url https://api.example.com/api/check_restricts
```


```
{ "country": "<string>", "ip": "<string>", "any_restricted": true, "platforms": { "polymarket": { "blocked": true, "close_only": true }, "opinion": { "blocked": true, "close_only": true }, "probable": { "blocked": true, "close_only": true }, "predictfun": { "blocked": true, "close_only": true }, "limitless": { "blocked": true, "close_only": true } } }
```


```
{ "country": "<string>", "ip": "<string>", "any_restricted": true, "platforms": { "polymarket": { "blocked": true, "close_only": true }, "opinion": { "blocked": true, "close_only": true }, "probable": { "blocked": true, "close_only": true }, "predictfun": { "blocked": true, "close_only": true }, "limitless": { "blocked": true, "close_only": true } } }
```


# Check regional availability
Returns whether trading is available for the current user region.
Use this endpoint to decide if trading actions should be enabled, disabled, or limited per platform.
cURL

```
curl --request GET \ --url https://api.example.com/api/check_restricts
```


```
curl --request GET \ --url https://api.example.com/api/check_restricts
```


```
{ "country": "<string>", "ip": "<string>", "any_restricted": true, "platforms": { "polymarket": { "blocked": true, "close_only": true }, "opinion": { "blocked": true, "close_only": true }, "probable": { "blocked": true, "close_only": true }, "predictfun": { "blocked": true, "close_only": true }, "limitless": { "blocked": true, "close_only": true } } }
```


```
{ "country": "<string>", "ip": "<string>", "any_restricted": true, "platforms": { "polymarket": { "blocked": true, "close_only": true }, "opinion": { "blocked": true, "close_only": true }, "probable": { "blocked": true, "close_only": true }, "predictfun": { "blocked": true, "close_only": true }, "limitless": { "blocked": true, "close_only": true } } }
```


## Documentation Index
Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/restricts/check-regional-availability#response-country)
Detected country code for the current request.
[​](https://docs.premarket.me/api-reference/restricts/check-regional-availability#response-ip)
Detected client IP address.
[​](https://docs.premarket.me/api-reference/restricts/check-regional-availability#response-any-restricted)
true if at least one platform has restrictions.

```
true
```

[​](https://docs.premarket.me/api-reference/restricts/check-regional-availability#response-platforms)
Restriction status by platform.
Show child attributes
[Get portfolio positions](https://docs.premarket.me/api-reference/positions/get-portfolio-positions)
[Get arbitrage opportunities](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

