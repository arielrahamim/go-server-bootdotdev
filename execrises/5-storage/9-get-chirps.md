# Get All Chirps

_Chapter: Storage | Slug: 9-get-chirps | UUID: 341b80d4-556f-4c5b-8afc-ffd12d5238c2_

# Get All Chirps

We need a way to retrieve _all_ the chirps from the database. Later, we'll add sorting and filtering functionality, but you can think of this as a very basic version of an endpoint that might serve a timeline of chirps.

## Assignment

1. [ ] Add a new query that retrieves all chirps in ascending order by `created_at`.
2. [ ] Add a `GET /api/chirps` endpoint that returns all chirps in the database. It should return them in the same structure as the `POST /api/chirps` endpoint, but as an array. Use a `200` status code for success. Order them by `created_at` in ascending order.

```json
[
  {
    "id": "94b7e44c-3604-42e3-bef7-ebfcc3efff8f",
    "created_at": "2021-01-01T00:00:00Z",
    "updated_at": "2021-01-01T00:00:00Z",
    "body": "Yo fam this feast is lit ong",
    "user_id": "123e4567-e89b-12d3-a456-426614174000"
  },
  {
    "id": "f0f87ec2-a8b5-48cc-b66a-a85ce7c7b862",
    "created_at": "2022-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z",
    "body": "What's good king?",
    "user_id": "123e4567-e89b-12d3-a456-426614174000"
  }
]
```

**Run and submit** the CLI tests.


## CLI

- Run: `bootdev run 341b80d4-556f-4c5b-8afc-ffd12d5238c2`
- Submit: `bootdev run 341b80d4-556f-4c5b-8afc-ffd12d5238c2 -s`
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
  "body": "If you're committed enough, you can make any story work.",
  "user_id": "${userID1}"
}
```
- Expect status: 201

### Step 4
- **POST** `${baseURL}/api/chirps`
- Body:
```json
{
  "body": "I once told a woman I was Kevin Costner, and it worked because I believed it.",
  "user_id": "${userID1}"
}
```
- Expect status: 201

### Step 5
- **GET** `${baseURL}/api/chirps`
- Expect status: 200
- JSON `.[0].body` eq `If you're committed enough, you can make any story work.`
- JSON `.[1].body` eq `I once told a woman I was Kevin Costner, and it worked because I believed it.`

