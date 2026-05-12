Title: Refresh session - Premarket Docs

Description: Issues a new access token for an already signed-in user. Use this endpoint when your current access token expires, while your refresh token is still valid.

Source: https://docs.premarket.me/api-reference/auth/refresh-session

---

[Skip to main content](https://docs.premarket.me/api-reference/auth/refresh-session#content-area)
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
curl --request POST \ --url https://api.example.com/api/auth/refresh \ --header 'Content-Type: application/json' \ --data ' { "refresh_token": "<string>" } '
```


```
curl --request POST \ --url https://api.example.com/api/auth/refresh \ --header 'Content-Type: application/json' \ --data ' { "refresh_token": "<string>" } '
```


```
{ "access_token": "<string>" }
```


```
{ "access_token": "<string>" }
```


# Refresh session
Issues a new access token for an already signed-in user.
Use this endpoint when your current access token expires, while your refresh token is still valid.
cURL

```
curl --request POST \ --url https://api.example.com/api/auth/refresh \ --header 'Content-Type: application/json' \ --data ' { "refresh_token": "<string>" } '
```


```
curl --request POST \ --url https://api.example.com/api/auth/refresh \ --header 'Content-Type: application/json' \ --data ' { "refresh_token": "<string>" } '
```


```
{ "access_token": "<string>" }
```


```
{ "access_token": "<string>" }
```


## Documentation Index
Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Body
[​](https://docs.premarket.me/api-reference/auth/refresh-session#body-refresh-token)

#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/auth/refresh-session#response-access-token)
[Sign in with wallet](https://docs.premarket.me/api-reference/auth/sign-in-with-wallet)
[Create API key](https://docs.premarket.me/api-reference/auth/create-api-key)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

