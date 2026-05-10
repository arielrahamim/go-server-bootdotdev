# Get Chirp

_Chapter: Storage | Slug: 10-get-chirp | UUID: 0a07a4a3-c11f-429f-ac70-52fa2e016bc0_

# Get Chirp

Now we need a way to lookup a single chirp by its ID. You might be thinking:

> "If I can get all of the chirps, why do I need a way to get just one?"

Imagine there are 10,000 chirps in the database - no, imagine 10,000,000,000! We'll obviously need to change our `GET /api/chirps` endpoint to only return a subset of chirps at a time.

However, our users will still need a way to view a single chirp - for example, maybe they have a link directly to it.

## Assignment

1. [ ] Add a `GET /api/chirps/{chirpID}` endpoint that returns a single chirp by its ID. The chirp ID will be passed in as a path parameter. For example:

```
GET /api/chirps/94b7e44c-3604-42e3-bef7-ebfcc3efff8f
```

You can get the string value of the path parameter like in Go with the [`http.Request.PathValue`](https://pkg.go.dev/net/http#Request.PathValue) method.

2. [ ] If the chirp is found, return it like so with a `200` code:

```json
{
  "id": "94b7e44c-3604-42e3-bef7-ebfcc3efff8f",
  "created_at": "2021-01-01T00:00:00Z",
  "updated_at": "2021-01-01T00:00:00Z",
  "body": "fr? no clowning?",
  "user_id": "123e4567-e89b-12d3-a456-426614174000"
}
```

3. [ ] Otherwise, return a `404`.

**Run and submit** the CLI tests.


## CLI

- Run: `bootdev run 0a07a4a3-c11f-429f-ac70-52fa2e016bc0`
- Submit: `bootdev run 0a07a4a3-c11f-429f-ac70-52fa2e016bc0 -s`
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
  "email": "saul@bettercall.com"
}
```
- Expect status: 201
- Capture variable `userID1` from `.id`

### Step 3
- **POST** `${baseURL}/api/chirps`
- Body:
```json
{
  "body": "I'm gonna be a damn good developer, and people are gonna know about it.",
  "user_id": "${userID1}"
}
```
- Expect status: 201
- Capture variable `chirpID1` from `.id`

### Step 4
- **GET** `${baseURL}/api/chirps/${chirpID1}`
- Expect status: 200
- JSON `.body` eq `I'm gonna be a damn good developer, and people are gonna know about it.`

### Step 5
- **GET** `${baseURL}/api/chirps/123e4567-e89b-12d3-a456-426614174059`
- Expect status: 404

