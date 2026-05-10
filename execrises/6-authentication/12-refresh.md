# Refresh Tokens

_Chapter: Authentication | Slug: 12-refresh | UUID: f7285cef-5185-4b15-b5fc-9533ccaafe8a_

# Refresh Tokens

To allow our users to stay logged in for longer periods, let's add refresh tokens to our authentication system. At the same time, we'll reduce the lifespan of our access tokens to improve security.

## Session Store

In our case, a refresh token will just be a random 256-bit string. It's a _token_, but not a _JSON Web Token_. It doesn't need to be a JWT because we'll store it in our database and associate it with a user server-side. No point in using stateless JWTs if we're going to store them in a database anyway.

To revoke a refresh token, we'll set a `revoked_at` timestamp in the database. If `revoked_at` is not null, the token is revoked and will be considered invalid.

## Assignment

1. [ ] Create a new database table with up/down migrations called `refresh_tokens`.
   - `token`: the primary key - it's just a string
   - `created_at`
   - `updated_at`
   - `user_id`: foreign key that deletes the row if the user is deleted
   - `expires_at`: the timestamp when the token expires
   - `revoked_at`: the timestamp when the token was revoked (null if not revoked)
2. [ ] Add a `func MakeRefreshToken() string` function to your `internal/auth` package. It should use the following to generate a random 256-bit (32-byte) hex-encoded string:
   - [rand.Read](https://pkg.go.dev/crypto/rand#Read) to generate 32 bytes (256 bits) of random data from the `crypto/rand` package (`math/rand`'s `Read` function is deprecated).
   - [hex.EncodeToString](https://pkg.go.dev/encoding/hex#EncodeToString) to convert the random data to a hex string
3. [ ] Update the `POST /api/login` endpoint to return a refresh token, as well as an access token:
   - [ ] Access tokens (JWTs) should expire after 1 hour. Expiration time is stored in the `exp` claim. You can remove the optional `expires_in_seconds` parameter from the endpoint.
   - [ ] Refresh tokens should expire after 60 days. Expiration time is stored in the database.
   - [ ] The `revoked_at` field should be null when the token is created.

```json
{
  "id": "5a47789c-a617-444a-8a80-b50359247804",
  "created_at": "2021-07-01T00:00:00Z",
  "updated_at": "2021-07-01T00:00:00Z",
  "email": "lane@example.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
  "refresh_token": "56aa826d22baab4b5ec2cea41a59ecbba03e542aedbb31d9b80326ac8ffcfa2a"
}
```

4. [ ] Create a `POST /api/refresh` endpoint. This new endpoint does _not_ accept a request body, but _does_ require a **refresh token** to be present in the headers, in the same `Authorization: Bearer <refresh-token>` format.

Look up the refresh token in the database. If it doesn't exist, or if it's expired or revoked, respond with a `401` status code. Otherwise, respond with a `200` code and this shape:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

The `token` field should be a newly created access token _for the given user_ that expires in 1 hour. I wrote a `GetUserFromRefreshToken` SQL query.

5. [ ] Create a new `POST /api/revoke` endpoint. This new endpoint does _not_ accept a request body, but _does_ require a **refresh token** to be present in the headers, in the same `Authorization: Bearer <refresh-token>` format.

Revoke the refresh token record in the database that matches the refresh token passed in the request header by setting the `revoked_at` to the current timestamp. Remember that any time you update a record, you should also be updating the `updated_at` timestamp.

Respond with a [`204` status code](https://www.rfc-editor.org/rfc/rfc9110.html#name-204-no-content). A 204 status means the request was successful but no body is returned.

**Run and submit** the CLI tests.


## CLI

- Run: `bootdev run f7285cef-5185-4b15-b5fc-9533ccaafe8a`
- Submit: `bootdev run f7285cef-5185-4b15-b5fc-9533ccaafe8a -s`
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
  "email": "saul@bettercall.com",
  "password": "123456"
}
```
- Expect status: 201
- JSON `.email` eq `saul@bettercall.com`

### Step 3
- **POST** `${baseURL}/api/login`
- Body:
```json
{
  "email": "saul@bettercall.com",
  "password": "123456"
}
```
- Expect status: 200
- Capture variable `saulAccessToken` from `.token`
- Capture variable `saulRefreshToken` from `.refresh_token`

### Step 4
- **POST** `${baseURL}/api/chirps`
- Headers: `{"Authorization": "Bearer ${saulRefreshToken}"}`
- Body:
```json
{
  "body": "Let's just say I know a guy... who knows a guy... who knows another guy."
}
```
- Expect status: 401

### Step 5
- **POST** `${baseURL}/api/chirps`
- Headers: `{"Authorization": "Bearer ${saulAccessToken}"}`
- Body:
```json
{
  "body": "Let's just say I know a guy... who knows a guy... who knows another guy."
}
```
- Expect status: 201

### Step 6
- **POST** `${baseURL}/api/refresh`
- Headers: `{"Authorization": "Bearer ${saulRefreshToken}"}`
- Expect status: 200
- Capture variable `saulAccessToken2` from `.token`

### Step 7
- **POST** `${baseURL}/api/chirps`
- Headers: `{"Authorization": "Bearer ${saulAccessToken2}"}`
- Body:
```json
{
  "body": "I'm the guy who's gonna win you this case."
}
```
- Expect status: 201

### Step 8
- **POST** `${baseURL}/api/revoke`
- Headers: `{"Authorization": "Bearer ${saulRefreshToken}"}`
- Expect status: 204

### Step 9
- **POST** `${baseURL}/api/refresh`
- Headers: `{"Authorization": "Bearer ${saulRefreshToken}"}`
- Expect status: 401

