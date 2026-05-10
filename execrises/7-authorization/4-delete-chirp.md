# Delete Chirp

_Chapter: Authorization | Slug: 4-delete-chirp | UUID: 61628ee7-a227-45a2-ab79-2721a52db32a_

# Delete Chirp

Oh no... the Chirpy CEO is chirping again. He's about to get the entire company cancelled. Let's add delete functionality!

## Assignment

1. [ ] Add a new `DELETE /api/chirps/{chirpID}` route to your server that deletes a chirp from the database by its `id`.
   - [ ] This is an authenticated endpoint, so be sure to check the token in the header. Only allow the deletion of a chirp if the user is the author of the chirp.
   - [ ] If they are not, return a `403` status code.
2. [ ] If the chirp is deleted successfully, return a `204` status code.
3. [ ] If the chirp is not found, return a `404` status code.

**Run and submit** the CLI tests.


## CLI

- Run: `bootdev run 61628ee7-a227-45a2-ab79-2721a52db32a`
- Submit: `bootdev run 61628ee7-a227-45a2-ab79-2721a52db32a -s`
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
- Capture variable `walterAccessToken` from `.token`

### Step 4
- **POST** `${baseURL}/api/chirps`
- Headers: `{"Authorization": "Bearer ${walterAccessToken}"}`
- Body:
```json
{
  "body": "I did it for me. I liked it. I was good at it. And I was really... I was alive."
}
```
- Expect status: 201
- Capture variable `chirpID` from `.id`

### Step 5
- **GET** `${baseURL}/api/chirps/${chirpID}`
- Expect status: 200

### Step 6
- **DELETE** `${baseURL}/api/chirps/${chirpID}`
- Expect status: 401

### Step 7
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

### Step 8
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

### Step 9
- **DELETE** `${baseURL}/api/chirps/${chirpID}`
- Headers: `{"Authorization": "Bearer ${saulAccessToken}"}`
- Expect status: 403

### Step 10
- **DELETE** `${baseURL}/api/chirps/${chirpID}`
- Headers: `{"Authorization": "Bearer ${walterAccessToken}"}`
- Expect status: 204

### Step 11
- **GET** `${baseURL}/api/chirps/${chirpID}`
- Expect status: 404

