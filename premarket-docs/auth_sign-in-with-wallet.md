Title: Sign in with wallet - Premarket Docs

Description: Signs in a user with a wallet signature and starts an authenticated session. If signature verification succeeds, the response returns: - access_token for authorized API requests; - refresh_token to renew the session later.

Source: https://docs.premarket.me/api-reference/auth/sign-in-with-wallet

---

[Skip to main content](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet#content-area)
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
- auth[POSTSign in with wallet](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet)[POSTRefresh session](https://docs.premarket.me/api-reference/auth/refresh-session)[POSTCreate API key](https://docs.premarket.me/api-reference/auth/create-api-key)
- [POSTSign in with wallet](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet)
- [POSTRefresh session](https://docs.premarket.me/api-reference/auth/refresh-session)
- [POSTCreate API key](https://docs.premarket.me/api-reference/auth/create-api-key)
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
- [POSTSign in with wallet](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet)
- [POSTRefresh session](https://docs.premarket.me/api-reference/auth/refresh-session)
- [POSTCreate API key](https://docs.premarket.me/api-reference/auth/create-api-key)
[POSTSign in with wallet](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet)
[POSTRefresh session](https://docs.premarket.me/api-reference/auth/refresh-session)
[POSTCreate API key](https://docs.premarket.me/api-reference/auth/create-api-key)
cURL

```
curl --request POST \ --url https://api.example.com/api/auth \ --header 'Content-Type: application/json' \ --data ' { "address": "<string>", "chain_id": 123, "signature": "<string>" } '
```


```
curl --request POST \ --url https://api.example.com/api/auth \ --header 'Content-Type: application/json' \ --data ' { "address": "<string>", "chain_id": 123, "signature": "<string>" } '
```


```
{ "access_token": "<string>", "refresh_token": "<string>" }
```


```
{ "access_token": "<string>", "refresh_token": "<string>" }
```


# Sign in with wallet
Signs in a user with a wallet signature and starts an authenticated session.
If signature verification succeeds, the response returns:
- access_token for authorized API requests;
- refresh_token to renew the session later.

```
access_token
```


```
refresh_token
```

cURL

```
curl --request POST \ --url https://api.example.com/api/auth \ --header 'Content-Type: application/json' \ --data ' { "address": "<string>", "chain_id": 123, "signature": "<string>" } '
```


```
curl --request POST \ --url https://api.example.com/api/auth \ --header 'Content-Type: application/json' \ --data ' { "address": "<string>", "chain_id": 123, "signature": "<string>" } '
```


```
{ "access_token": "<string>", "refresh_token": "<string>" }
```


```
{ "access_token": "<string>", "refresh_token": "<string>" }
```


## Documentation Index
Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Body
[​](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet#body-address)

```
10 - 100
```

[​](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet#body-chain-id)
[​](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet#body-signature)

```
10 - 200
```


#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet#response-access-token)
[​](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet#response-refresh-token)
[Get event details](https://docs.premarket.me/api-reference/events/get-event-details)
[Refresh session](https://docs.premarket.me/api-reference/auth/refresh-session)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

