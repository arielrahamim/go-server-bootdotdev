# Sorting Chirps

_Chapter: Documentation | Slug: 4-sorting-chirp | UUID: 2f20da66-64d8-47b4-8678-4a95cd06767a_

# Sorting Chirps

A common feature in APIs is the ability to sort the response by a field. We don't want to add additional endpoints for every possible sort order, so we'll use a query parameter instead.

## Assignment

Update the `GET /api/chirps` endpoint. It should accept an _optional_ query parameter called `sort`. It can have 2 possible values:

- `asc` - Sort the chirps in the response by `created_at` in ascending order
- `desc` - Sort the chirps in the response by `created_at` in descending order

`asc` is the default if no `sort` query parameter is provided.

> [!tip]
> Keep it simple! You can just sort the chirps in-memory using [`sort.Slice`](https://pkg.go.dev/sort#Slice).

**Run and submit** the CLI tests.

## Examples of Valid URLs

- `GET http://localhost:8080/api/chirps?sort=asc`
- `GET http://localhost:8080/api/chirps?sort=desc`
- `GET http://localhost:8080/api/chirps`


## CLI

- Run: `bootdev run 2f20da66-64d8-47b4-8678-4a95cd06767a`
- Submit: `bootdev run 2f20da66-64d8-47b4-8678-4a95cd06767a -s`
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
- **POST** `${baseURL}/api/chirps`
- Headers: `{"Authorization": "Bearer ${waltAccessToken}"}`
- Body:
```json
{
  "body": "Cmon Pinkman"
}
```
- Expect status: 201
- JSON `.body` eq `Cmon Pinkman`

### Step 7
- **POST** `${baseURL}/api/chirps`
- Headers: `{"Authorization": "Bearer ${waltAccessToken}"}`
- Body:
```json
{
  "body": "Darn that fly, I just wanna cook"
}
```
- Expect status: 201
- JSON `.body` eq `Darn that fly, I just wanna cook`

### Step 8
- **GET** `${baseURL}/api/chirps?sort=desc`
- Expect status: 200
- JSON `.[0].body` eq `Darn that fly, I just wanna cook`
- JSON `.[1].body` eq `Cmon Pinkman`
- JSON `.[2].body` eq `Gale!`
- JSON `.[3].body` eq `I'm the one who knocks!`

### Step 9
- **GET** `${baseURL}/api/chirps?sort=asc`
- Expect status: 200
- JSON `.[0].body` eq `I'm the one who knocks!`
- JSON `.[1].body` eq `Gale!`
- JSON `.[2].body` eq `Cmon Pinkman`
- JSON `.[3].body` eq `Darn that fly, I just wanna cook`

