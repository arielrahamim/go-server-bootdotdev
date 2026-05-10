# API Keys

_Chapter: Webhooks | Slug: 4-apikey | UUID: eb86a031-f0f3-4939-9221-7112e5d272cf_

# API Keys

You may have noticed that there is an issue with our webhook handler: it's not secure!

Anyone can send a request to our webhook handler, and we'll process it. That means that if Chirpy users figured out our API documentation, they could simply upgrade their account without paying!

## Assignment

Luckily, Polka has a solution for this: API keys. Polka provided us with an API key, and if a request to our webhook handler doesn't use that API key, we should reject the request. This ensures that only Polka can tell us to upgrade a user's account.

Your Polka key: `f271c81ff7084ee5b99a5091b42d486e`

1. [ ] Add a new secret value to your `.env` file called `POLKA_KEY`. This is the api key that polka will send so that we know it's them (and not someone else trying to get free Chirpy red). Load it into your server and store it in your `apiConfig`.
2. [ ] Add a `func GetAPIKey(headers http.Header) (string, error)` to your `auth` package. It should extract the api key from the `Authorization` header, which is expected to be in this format:

```
Authorization: ApiKey THE_KEY_HERE
```

You'll need to strip out the `ApiKey` part and the whitespace and return just the key.

3. [ ] Update the `POST /api/polka/webhooks` endpoint. It should ensure that the API key in the header matches the one stored in the `.env` file. If it doesn't, the endpoint should respond with a `401` status code.

**Run and submit** the CLI tests.


## CLI

- Run: `bootdev run eb86a031-f0f3-4939-9221-7112e5d272cf`
- Submit: `bootdev run eb86a031-f0f3-4939-9221-7112e5d272cf -s`
- Default base URL: `http://localhost:8080`

## Test Steps

### Step 1
- **POST** `${baseURL}/admin/reset`
- Expect status: 200

### Step 2
- **POST** `${baseURL}/api/users`
- Body:
```json
{
  "email": "walt@breakingbad.com",
  "password": "123456"
}
```
- Expect status: 201
- JSON `.email` eq `walt@breakingbad.com`
- JSON `.is_chirpy_red` eq `None`
- Capture variable `userID` from `.id`

### Step 3
- **POST** `${baseURL}/api/polka/webhooks`
- Body:
```json
{
  "data": {
    "user_id": "${userID}"
  },
  "event": "user.upgraded"
}
```
- Expect status: 401

### Step 4
- **POST** `${baseURL}/api/polka/webhooks`
- Headers: `{"Authorization": "ApiKey f271c81ff7084ee5b99a5091b42d486e"}`
- Body:
```json
{
  "data": {
    "user_id": "${userID}"
  },
  "event": "user.upgraded"
}
```
- Expect status: 204

### Step 5
- **POST** `${baseURL}/api/login`
- Body:
```json
{
  "email": "walt@breakingbad.com",
  "password": "123456"
}
```
- Expect status: 200
- JSON `.email` eq `walt@breakingbad.com`
- JSON `.is_chirpy_red` eq `None`

