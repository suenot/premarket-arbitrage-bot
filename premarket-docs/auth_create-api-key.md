Title: Create API key - Premarket Docs

Description: Creates a new API key for your account. Use it when you want to access API-key protected endpoints from scripts, bots, or external tools. Authorization: - Requires Authorization: Bearer . - Available only for signed-in users.

Source: https://docs.premarket.me/api-reference/auth/create-api-key

---

[Skip to main content](https://docs.premarket.me/api-reference/auth/create-api-key#content-area)
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
curl --request POST \ --url https://api.example.com/api/api-key/generate \ --header 'Authorization: Bearer <token>'
```


```
curl --request POST \ --url https://api.example.com/api/api-key/generate \ --header 'Authorization: Bearer <token>'
```


```
{ "api_key": "<string>" }
```


```
{ "api_key": "<string>" }
```


# Create API key
Creates a new API key for your account.
Use it when you want to access API-key protected endpoints from scripts, bots, or external tools.
Authorization:
- Requires Authorization: Bearer <access_token>.
- Available only for signed-in users.

```
Authorization: Bearer <access_token>
```

cURL

```
curl --request POST \ --url https://api.example.com/api/api-key/generate \ --header 'Authorization: Bearer <token>'
```


```
curl --request POST \ --url https://api.example.com/api/api-key/generate \ --header 'Authorization: Bearer <token>'
```


```
{ "api_key": "<string>" }
```


```
{ "api_key": "<string>" }
```


## Documentation Index
Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Authorizations
[​](https://docs.premarket.me/api-reference/auth/create-api-key#authorization-authorization)
Bearer authentication header of the form Bearer <token>, where <token> is your auth token.

```
Bearer <token>
```


```
<token>
```


#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/auth/create-api-key#response-api-key)
[Refresh session](https://docs.premarket.me/api-reference/auth/refresh-session)
[Get wallet cash balances](https://docs.premarket.me/api-reference/balances/get-wallet-cash-balances)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

