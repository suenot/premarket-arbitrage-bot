Title: Get event details - Premarket Docs

Description: Returns full details for one event by event_slug. Use this endpoint to open an event page and get: - event title, icon, slug, end date and data sources; - outcome list with current prices across available platforms; - volume and transactions summary; - per-outcome market ids and tokens that can be reused in /api/market/{market_id} and /api/orderbook. slug is the event identifier used in event links and in responses of list endpoints.

Source: https://docs.premarket.me/api-reference/events/get-event-details

---

[Skip to main content](https://docs.premarket.me/api-reference/events/get-event-details#content-area)
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
curl --request GET \ --url https://api.example.com/api/event/{slug}
```


```
curl --request GET \ --url https://api.example.com/api/event/{slug}
```


```
{ "event_id": "<string>", "primary_source": "<string>", "event_title": "<string>", "event_icon": "<string>", "event_slug": "<string>", "end_date": "<string>", "is_closed": true, "sources": [ "<string>" ], "descriptions": {}, "volume": 123, "transactions": 123, "outcomes": [ { "outcome_group_title": "<string>", "icon": "<string>", "best_yes_price": 123, "best_yes_source": "<string>", "best_no_price": 123, "best_no_source": "<string>", "end_date": "<string>", "is_active": false, "market_status": "<string>", "resolved_result": "<string>", "closed_reason": "<string>", "is_date_market": true, "market_date": "<string>", "sort_ts": 123, "polymarket_yes": 123, "polymarket_yes_token": "<string>", "polymarket_no": 123, "polymarket_no_token": "<string>", "polymarket_market_id": "<string>", "polymarket_condition_id": "<string>", "polymarket_is_active": true, "opinion_yes": 123, "opinion_yes_token": "<string>", "opinion_no": 123, "opinion_no_token": "<string>", "opinion_market_id": "<string>", "opinion_condition_id": "<string>", "opinion_is_active": true, "probable_yes": 123, "probable_yes_token": "<string>", "probable_no": 123, "probable_no_token": "<string>", "probable_market_id": "<string>", "probable_condition_id": "<string>", "probable_is_active": true, "predictfun_yes": 123, "predictfun_yes_token": "<string>", "predictfun_no": 123, "predictfun_no_token": "<string>", "predictfun_market_id": "<string>", "predictfun_condition_id": "<string>", "predictfun_is_active": true, "limitless_yes": 123, "limitless_yes_token": "<string>", "limitless_no": 123, "limitless_no_token": "<string>", "limitless_market_id": "<string>", "limitless_condition_id": "<string>", "limitless_is_active": true, "cached_orderbook": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes_price": 123, "no_price": 123, "yes": { "asks": [ [ 123, 123 ] ], "bids": [ [ 123, 123 ] ] }, "no": { "asks": [ [ 123, 123 ] ], "bids": [ [ 123, 123 ] ] } } } ] }
```

```
{ "event_id": "<string>", "primary_source": "<string>", "event_title": "<string>", "event_icon": "<string>", "event_slug": "<string>", "end_date": "<string>", "is_closed": true, "sources": [ "<string>" ], "descriptions": {}, "volume": 123, "transactions": 123, "outcomes": [ { "outcome_group_title": "<string>", "icon": "<string>", "best_yes_price": 123, "best_yes_source": "<string>", "best_no_price": 123, "best_no_source": "<string>", "end_date": "<string>", "is_active": false, "market_status": "<string>", "resolved_result": "<string>", "closed_reason": "<string>", "is_date_market": true, "market_date": "<string>", "sort_ts": 123, "polymarket_yes": 123, "polymarket_yes_token": "<string>", "polymarket_no": 123, "polymarket_no_token": "<string>", "polymarket_market_id": "<string>", "polymarket_condition_id": "<string>", "polymarket_is_active": true, "opinion_yes": 123, "opinion_yes_token": "<string>", "opinion_no": 123, "opinion_no_token": "<string>", "opinion_market_id": "<string>", "opinion_condition_id": "<string>", "opinion_is_active": true, "probable_yes": 123, "probable_yes_token": "<string>", "probable_no": 123, "probable_no_token": "<string>", "probable_market_id": "<string>", "probable_condition_id": "<string>", "probable_is_active": true, "predictfun_yes": 123, "predictfun_yes_token": "<string>", "predictfun_no": 123, "predictfun_no_token": "<string>", "predictfun_market_id": "<string>", "predictfun_condition_id": "<string>", "predictfun_is_active": true, "limitless_yes": 123, "limitless_yes_token": "<string>", "limitless_no": 123, "limitless_no_token": "<string>", "limitless_market_id": "<string>", "limitless_condition_id": "<string>", "limitless_is_active": true, "cached_orderbook": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes_price": 123, "no_price": 123, "yes": { "asks": [ [ 123, 123 ] ], "bids": [ [ 123, 123 ] ] }, "no": { "asks": [ [ 123, 123 ] ], "bids": [ [ 123, 123 ] ] } } } ] }
```

Returns full details for one event by event_slug.

```
event_slug
```

Use this endpoint to open an event page and get:
- event title, icon, slug, end date and data sources;
- outcome list with current prices across available platforms;
- volume and transactions summary;
- per-outcome market ids and tokens that can be reused in /api/market/{market_id} and /api/orderbook.

```
/api/market/{market_id}
```


```
/api/orderbook
```

slug is the event identifier used in event links and in responses of list endpoints.

```
slug
```

cURL

```
curl --request GET \ --url https://api.example.com/api/event/{slug}
```


```
curl --request GET \ --url https://api.example.com/api/event/{slug}
```


```
{ "event_id": "<string>", "primary_source": "<string>", "event_title": "<string>", "event_icon": "<string>", "event_slug": "<string>", "end_date": "<string>", "is_closed": true, "sources": [ "<string>" ], "descriptions": {}, "volume": 123, "transactions": 123, "outcomes": [ { "outcome_group_title": "<string>", "icon": "<string>", "best_yes_price": 123, "best_yes_source": "<string>", "best_no_price": 123, "best_no_source": "<string>", "end_date": "<string>", "is_active": false, "market_status": "<string>", "resolved_result": "<string>", "closed_reason": "<string>", "is_date_market": true, "market_date": "<string>", "sort_ts": 123, "polymarket_yes": 123, "polymarket_yes_token": "<string>", "polymarket_no": 123, "polymarket_no_token": "<string>", "polymarket_market_id": "<string>", "polymarket_condition_id": "<string>", "polymarket_is_active": true, "opinion_yes": 123, "opinion_yes_token": "<string>", "opinion_no": 123, "opinion_no_token": "<string>", "opinion_market_id": "<string>", "opinion_condition_id": "<string>", "opinion_is_active": true, "probable_yes": 123, "probable_yes_token": "<string>", "probable_no": 123, "probable_no_token": "<string>", "probable_market_id": "<string>", "probable_condition_id": "<string>", "probable_is_active": true, "predictfun_yes": 123, "predictfun_yes_token": "<string>", "predictfun_no": 123, "predictfun_no_token": "<string>", "predictfun_market_id": "<string>", "predictfun_condition_id": "<string>", "predictfun_is_active": true, "limitless_yes": 123, "limitless_yes_token": "<string>", "limitless_no": 123, "limitless_no_token": "<string>", "limitless_market_id": "<string>", "limitless_condition_id": "<string>", "limitless_is_active": true, "cached_orderbook": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes_price": 123, "no_price": 123, "yes": { "asks": [ [ 123, 123 ] ], "bids": [ [ 123, 123 ] ] }, "no": { "asks": [ [ 123, 123 ] ], "bids": [ [ 123, 123 ] ] } } } ] }
```

```
{ "event_id": "<string>", "primary_source": "<string>", "event_title": "<string>", "event_icon": "<string>", "event_slug": "<string>", "end_date": "<string>", "is_closed": true, "sources": [ "<string>" ], "descriptions": {}, "volume": 123, "transactions": 123, "outcomes": [ { "outcome_group_title": "<string>", "icon": "<string>", "best_yes_price": 123, "best_yes_source": "<string>", "best_no_price": 123, "best_no_source": "<string>", "end_date": "<string>", "is_active": false, "market_status": "<string>", "resolved_result": "<string>", "closed_reason": "<string>", "is_date_market": true, "market_date": "<string>", "sort_ts": 123, "polymarket_yes": 123, "polymarket_yes_token": "<string>", "polymarket_no": 123, "polymarket_no_token": "<string>", "polymarket_market_id": "<string>", "polymarket_condition_id": "<string>", "polymarket_is_active": true, "opinion_yes": 123, "opinion_yes_token": "<string>", "opinion_no": 123, "opinion_no_token": "<string>", "opinion_market_id": "<string>", "opinion_condition_id": "<string>", "opinion_is_active": true, "probable_yes": 123, "probable_yes_token": "<string>", "probable_no": 123, "probable_no_token": "<string>", "probable_market_id": "<string>", "probable_condition_id": "<string>", "probable_is_active": true, "predictfun_yes": 123, "predictfun_yes_token": "<string>", "predictfun_no": 123, "predictfun_no_token": "<string>", "predictfun_market_id": "<string>", "predictfun_condition_id": "<string>", "predictfun_is_active": true, "limitless_yes": 123, "limitless_yes_token": "<string>", "limitless_no": 123, "limitless_no_token": "<string>", "limitless_market_id": "<string>", "limitless_condition_id": "<string>", "limitless_is_active": true, "cached_orderbook": { "yes_ask": 123, "yes_bid": 123, "no_ask": 123, "no_bid": 123, "yes_price": 123, "no_price": 123, "yes": { "asks": [ [ 123, 123 ] ], "bids": [ [ 123, 123 ] ] }, "no": { "asks": [ [ 123, 123 ] ], "bids": [ [ 123, 123 ] ] } } } ] }
```

Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Path Parameters
[​](https://docs.premarket.me/api-reference/events/get-event-details#parameter-slug)

#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-event-id)
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-primary-source)
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-event-title-one-of-0)
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-event-icon-one-of-0)
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-event-slug)
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-end-date-one-of-0)
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-is-closed)
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-sources)
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-descriptions)
Show child attributes
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-volume-one-of-0)
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-transactions-one-of-0)
[​](https://docs.premarket.me/api-reference/events/get-event-details#response-outcomes)
Show child attributes
[List market events](https://docs.premarket.me/api-reference/events/list-market-events)
[Sign in with wallet](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

