Title: Getting Started - Premarket Docs

Description: How to work with Premarket API

Source: https://docs.premarket.me/getting-started

---

[Skip to main content](https://docs.premarket.me/getting-started#content-area)
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
- v2/market
- [Base URL](https://docs.premarket.me/getting-started#base-url)
- [First Request](https://docs.premarket.me/getting-started#first-request)
- [Common Response Shape](https://docs.premarket.me/getting-started#common-response-shape)
- [Error Handling](https://docs.premarket.me/getting-started#error-handling)
- [Auth Notes](https://docs.premarket.me/getting-started#auth-notes)
[Base URL](https://docs.premarket.me/getting-started#base-url)
[First Request](https://docs.premarket.me/getting-started#first-request)
[Common Response Shape](https://docs.premarket.me/getting-started#common-response-shape)
[Error Handling](https://docs.premarket.me/getting-started#error-handling)
[Auth Notes](https://docs.premarket.me/getting-started#auth-notes)

# Getting Started
How to work with Premarket API

## Documentation Index
Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

## [​](https://docs.premarket.me/getting-started#base-url)Base URL
[​](https://docs.premarket.me/getting-started#base-url)

```
https://api.premarket.me
```


## [​](https://docs.premarket.me/getting-started#first-request)First Request
[​](https://docs.premarket.me/getting-started#first-request)

```
curl -sS "https://api.premarket.me/api/events/list?offset=0&limit=20"
```


```
curl -sS "https://api.premarket.me/api/events/list?offset=0&limit=20"
```


## [​](https://docs.premarket.me/getting-started#common-response-shape)Common Response Shape
[​](https://docs.premarket.me/getting-started#common-response-shape)
- count: number of returned items
- total: total items
- offset, limit: pagination controls

```
count
```


```
total
```


```
offset
```


```
limit
```


## [​](https://docs.premarket.me/getting-started#error-handling)Error Handling
[​](https://docs.premarket.me/getting-started#error-handling)
- 4xx: request/auth/validation issue
- 5xx: server or upstream issue

```
4xx
```


```
5xx
```


## [​](https://docs.premarket.me/getting-started#auth-notes)Auth Notes
[​](https://docs.premarket.me/getting-started#auth-notes)
[Premarket Docs](https://docs.premarket.me/)
[List market events](https://docs.premarket.me/api-reference/events/list-market-events)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

