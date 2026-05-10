# Documentation

_Chapter: Documentation | Slug: 1-filter-chirp | UUID: c1a4f8aa-de85-45fe-9e70-e49a98e14e3a_

# Documentation

When you're designing a server-side API, no one is going to know how to interact with it unless you tell them. Are you going to force the front-end developers, mobile developers, or other back-end service teams to sift through your code and reverse engineer your API?

Of course not! You're a good person. You're going to write documentation.

## First Be Obvious, Then Document It Anyway

We've talked a lot about how your REST API should follow conventions as much as possible. That said, the conventions _are not enough_. You still need to document your endpoints. Without documentation, no one will know:

- Which resources are available
- What the path to the endpoints are
- Which HTTP methods are supported for each resource
- What the shape of the data is for each resource
- etc.

## Assignment

One type of endpoint that's nearly impossible to interact with without documentation is a plural `GET` endpoint, that is, an endpoint that returns a list of resources. They often have different sorting, filtering, and [pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) features.

1. [ ] Update the `GET /api/chirps` endpoint. It should accept an _optional_ query parameter called `author_id`.
   - [ ] If the `author_id` query parameter is provided, the endpoint should return only the chirps for that author.
   - [ ] If the `author_id` query parameter is not provided, the endpoint should return all chirps as it did before.

For example:

`GET http://localhost:8080/api/chirps?author_id=1`

_Continue sorting the chirps by `created_at` in ascending order._

> [!warning]
> Be sure to filter by author ID at the database level, not in memory! That will be more efficient on large datasets.

**Run and submit** the CLI tests.

## Tips

The [http.Request](https://pkg.go.dev/net/http#Request) struct has a way to grab the query parameters from the URL:

```go
s := r.URL.Query().Get("author_id")
// s is a string that contains the value of the author_id query parameter
// if it exists, or an empty string if it doesn't
```


## CLI

- Run: `bootdev run c1a4f8aa-de85-45fe-9e70-e49a98e14e3a`
- Submit: `bootdev run c1a4f8aa-de85-45fe-9e70-e49a98e14e3a -s`
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
- Capture variable `waltID` from `.id`

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
- Capture variable `waltAccessToken` from `.token`

### Step 4
- **POST** `${baseURL}/api/chirps`
- Headers: `{"Authorization": "Bearer ${waltAccessToken}"}`
- Body:
```json
{
  "body": "I'm the one who knocks!"
}
```
- Expect status: 201
- JSON `.body` eq `I'm the one who knocks!`

### Step 5
- **POST** `${baseURL}/api/chirps`
- Headers: `{"Authorization": "Bearer ${waltAccessToken}"}`
- Body:
```json
{
  "body": "Gale!"
}
```
- Expect status: 201
- JSON `.body` eq `Gale!`

### Step 6
- **POST** `${baseURL}/api/users`
- Body:
```json
{
  "email": "skyler@breakingbad.com",
  "password": "000111"
}
```
- Expect status: 201
- JSON `.email` eq `skyler@breakingbad.com`
- Capture variable `skylerID` from `.id`

### Step 7
- **POST** `${baseURL}/api/login`
- Body:
```json
{
  "email": "skyler@breakingbad.com",
  "password": "000111"
}
```
- Expect status: 200
- JSON `.email` eq `skyler@breakingbad.com`
- Capture variable `skylerAccessToken` from `.token`

### Step 8
- **POST** `${baseURL}/api/chirps`
- Headers: `{"Authorization": "Bearer ${skylerAccessToken}"}`
- Body:
```json
{
  "body": "Mr President...."
}
```
- Expect status: 201
- JSON `.body` eq `Mr President....`

### Step 9
- **GET** `${baseURL}/api/chirps?author_id=${waltID}`
- Expect status: 200
- Body contains: `I'm the one who knocks!`
- Body contains: `Gale!`

### Step 10
- **GET** `${baseURL}/api/chirps?author_id=${skylerID}`
- Expect status: 200
- Body contains: `Mr President....`

