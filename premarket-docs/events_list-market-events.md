Title: List market events - Premarket Docs

Description: Returns the main list of market events with pagination and optional filters by platform and category. You can also change the order of results with the sort parameter.

Source: https://docs.premarket.me/api-reference/events/list-market-events

---

[Skip to main content](https://docs.premarket.me/api-reference/events/list-market-events#content-area)
[Premarket Docs home page](https://docs.premarket.me/)
- [X (Twitter)](https://x.com/pre_markets)
[X (Twitter)](https://x.com/pre_markets)

##### Overview
- [Premarket Docs](https://docs.premarket.me/)
- [Getting Started](https://docs.premarket.me/getting-started)
[Premarket Docs](https://docs.premarket.me/)
[Getting Started](https://docs.premarket.me/getting-started)

##### API Reference
- events[GETList market events](https://docs.premarket.me/api-reference/events/list-market-events)[GETGet event details](https://docs.premarket.me/api-reference/events/get-event-details)
- [GETList market events](https://docs.premarket.me/api-reference/events/list-market-events)
- [GETGet event details](https://docs.premarket.me/api-reference/events/get-event-details)
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
- [GETList market events](https://docs.premarket.me/api-reference/events/list-market-events)
- [GETGet event details](https://docs.premarket.me/api-reference/events/get-event-details)
[GETList market events](https://docs.premarket.me/api-reference/events/list-market-events)
[GETGet event details](https://docs.premarket.me/api-reference/events/get-event-details)
cURL

```
curl --request GET \ --url https://api.example.com/api/events/list
```


```
curl --request GET \ --url https://api.example.com/api/events/list
```


```
{ "events": [ { "event_id": "<string>", "event_title": "<string>", "icon": "<string>", "source": "<string>", "event_slug": "<string>", "market_count": 0, "is_paired": false, "sources": [ "<string>" ], "outcomes": [ { "name": "<string>", "chance": 123, "market_id": "<string>" } ], "transactions": 0, "volume": 0, "tags": [ "<string>" ], "category": "<string>" } ], "offset": 0, "limit": 100, "count": 0, "total": 0, "cached": false, "generated_at": "<string>" }
```


```
{ "events": [ { "event_id": "<string>", "event_title": "<string>", "icon": "<string>", "source": "<string>", "event_slug": "<string>", "market_count": 0, "is_paired": false, "sources": [ "<string>" ], "outcomes": [ { "name": "<string>", "chance": 123, "market_id": "<string>" } ], "transactions": 0, "volume": 0, "tags": [ "<string>" ], "category": "<string>" } ], "offset": 0, "limit": 100, "count": 0, "total": 0, "cached": false, "generated_at": "<string>" }
```

Returns the main list of market events with pagination and optional filters by platform and category. You can also change the order of results with the sort parameter.

```
sort
```

cURL

```
curl --request GET \ --url https://api.example.com/api/events/list
```


```
curl --request GET \ --url https://api.example.com/api/events/list
```


```
{ "events": [ { "event_id": "<string>", "event_title": "<string>", "icon": "<string>", "source": "<string>", "event_slug": "<string>", "market_count": 0, "is_paired": false, "sources": [ "<string>" ], "outcomes": [ { "name": "<string>", "chance": 123, "market_id": "<string>" } ], "transactions": 0, "volume": 0, "tags": [ "<string>" ], "category": "<string>" } ], "offset": 0, "limit": 100, "count": 0, "total": 0, "cached": false, "generated_at": "<string>" }
```


```
{ "events": [ { "event_id": "<string>", "event_title": "<string>", "icon": "<string>", "source": "<string>", "event_slug": "<string>", "market_count": 0, "is_paired": false, "sources": [ "<string>" ], "outcomes": [ { "name": "<string>", "chance": 123, "market_id": "<string>" } ], "transactions": 0, "volume": 0, "tags": [ "<string>" ], "category": "<string>" } ], "offset": 0, "limit": 100, "count": 0, "total": 0, "cached": false, "generated_at": "<string>" }
```

Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Query Parameters
[​](https://docs.premarket.me/api-reference/events/list-market-events#parameter-offset)
Zero-based offset in the filtered events list.

```
x >= 0
```

[​](https://docs.premarket.me/api-reference/events/list-market-events#parameter-limit)
Maximum number of events to return.

```
1 <= x <= 1000
```

[​](https://docs.premarket.me/api-reference/events/list-market-events#parameter-platforms)
Optional comma-separated platform ids, for example polymarket,opinion.

```
polymarket,opinion
```

[​](https://docs.premarket.me/api-reference/events/list-market-events#parameter-category)
Optional category slug used by the events feed filter.
[​](https://docs.premarket.me/api-reference/events/list-market-events#parameter-sort)
Optional sort mode. Use new to show newer events first or breaking to show the most active events first.

```
new
```


```
breaking
```


#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/events/list-market-events#response-events)
Show child attributes
[​](https://docs.premarket.me/api-reference/events/list-market-events#response-offset)
[​](https://docs.premarket.me/api-reference/events/list-market-events#response-limit)
[​](https://docs.premarket.me/api-reference/events/list-market-events#response-count)
[​](https://docs.premarket.me/api-reference/events/list-market-events#response-total)
[​](https://docs.premarket.me/api-reference/events/list-market-events#response-cached)
[​](https://docs.premarket.me/api-reference/events/list-market-events#response-generated-at-one-of-0)
[Getting Started](https://docs.premarket.me/getting-started)
[Get event details](https://docs.premarket.me/api-reference/events/get-event-details)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

