# Authorization

_Chapter: Authorization | Slug: 1-update-user | UUID: be14c814-e6c2-4b96-a361-e33bcfe71f00_

# Authorization

While authentication is about verifying _who_ a user is, authorization is about verifying _what a user is allowed to do_.

For example, a hypothetical YouTuber `ThePrimeagen` should be allowed to edit and delete the videos on his account, and everyone should be allowed to view them. Another absolutely-not-real YouTuber `TEEJ` should be able to view `ThePrimeagen`'s videos, but not edit or delete them.

Authorization logic is just the code that enforces these kinds of rules.

## Assignment

We already have a bit of authorization built into Chirpy: authenticated users can only create chirps for themselves, not for others.

1. [ ] Add a `PUT /api/users` endpoint so that users can update their own (but not others') email and password. It requires:
   - An access token in the header
   - A new `password` and `email` in the request body
2. [ ] Hash the password, then update the hashed password and the email for the authenticated user in the database. Respond with a `200` if everything is successful and the newly updated `User` resource (omitting the password of course).
3. [ ] If the access token is malformed or missing, respond with a `401` status code.

**Run and submit** the CLI tests.


## CLI

- Run: `bootdev run be14c814-e6c2-4b96-a361-e33bcfe71f00`
- Submit: `bootdev run be14c814-e6c2-4b96-a361-e33bcfe71f00 -s`
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

### Step 3
- **POST** `${baseURL}/api/login`
- Body:
```json
{
  "email": "walt@breakingbad.com",
  "password": "123456"
}
```
- Expect status: 200
- Capture variable `jwtToken` from `.token`

### Step 4
- **PUT** `${baseURL}/api/users`
- Headers: `{"Authorization": "Bearer ${jwtToken}"}`
- Body:
```json
{
  "email": "walter@breakingbad.com",
  "password": "losPollosHermanos"
}
```
- Expect status: 200
- JSON `.email` eq `walter@breakingbad.com`

### Step 5
- **PUT** `${baseURL}/api/users`
- Body:
```json
{
  "email": "walter@breakingbad.com",
  "password": "j3ssePinkM@nCantCook"
}
```
- Expect status: 401

### Step 6
- **PUT** `${baseURL}/api/users`
- Headers: `{"Authorization": "Bearer badToken"}`
- Body:
```json
{
  "email": "walter@breakingbad.com",
  "password": "j3ssePinkM@nCantCook"
}
```
- Expect status: 401

