Title: Get my profile - Premarket Docs

Description: Returns profile data for the signed-in user. Includes linked wallet address and current API key (if already generated).

Source: https://docs.premarket.me/api-reference/user/get-my-profile

---

[Skip to main content](https://docs.premarket.me/api-reference/user/get-my-profile#content-area)
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
- user[GETGet my profile](https://docs.premarket.me/api-reference/user/get-my-profile)
- [GETGet my profile](https://docs.premarket.me/api-reference/user/get-my-profile)
- v2/platforms
- v2/market
- [GETGet my profile](https://docs.premarket.me/api-reference/user/get-my-profile)
[GETGet my profile](https://docs.premarket.me/api-reference/user/get-my-profile)
cURL

```
curl --request GET \ --url https://api.example.com/api/profile \ --header 'Authorization: Bearer <token>'
```


```
curl --request GET \ --url https://api.example.com/api/profile \ --header 'Authorization: Bearer <token>'
```


```
{ "address": "<string>", "api_key": "<string>" }
```


```
{ "address": "<string>", "api_key": "<string>" }
```


# Get my profile
Returns profile data for the signed-in user.
Includes linked wallet address and current API key (if already generated).
cURL

```
curl --request GET \ --url https://api.example.com/api/profile \ --header 'Authorization: Bearer <token>'
```


```
curl --request GET \ --url https://api.example.com/api/profile \ --header 'Authorization: Bearer <token>'
```


```
{ "address": "<string>", "api_key": "<string>" }
```


```
{ "address": "<string>", "api_key": "<string>" }
```


## Documentation Index
Fetch the complete documentation index at: [https://docs.premarket.me/llms.txt](https://docs.premarket.me/llms.txt)
Use this file to discover all available pages before exploring further.

#### Authorizations
[​](https://docs.premarket.me/api-reference/user/get-my-profile#authorization-authorization)
Bearer authentication header of the form Bearer <token>, where <token> is your auth token.

```
Bearer <token>
```


```
<token>
```


#### Response
Successful Response
[​](https://docs.premarket.me/api-reference/user/get-my-profile#response-address-one-of-0)
[​](https://docs.premarket.me/api-reference/user/get-my-profile#response-api-key-one-of-0)
[Get arbitrage opportunities](https://docs.premarket.me/api-reference/arbitrage/get-arbitrage-opportunities)
[List platforms](https://docs.premarket.me/api-reference/v2platforms/list-platforms)
[Powered byThis documentation is built and hosted on Mintlify, a developer documentation platform](https://www.mintlify.com?utm_campaign=poweredBy&utm_medium=referral&utm_source=premarket-e1dc75dd)

