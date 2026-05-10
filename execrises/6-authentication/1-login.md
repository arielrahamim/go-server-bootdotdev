# Authentication With Passwords

_Chapter: Authentication | Slug: 1-login | UUID: 294e5c16-d1e8-4836-871c-dedc98581236_

# Authentication With Passwords

Authentication is the process of verifying _who_ a user is. If you don't have a secure authentication system, your back-end systems will be open to attack!

Imagine if I could make an HTTP request to the YouTube API and upload a video to _your_ channel. YouTube's authentication system prevents this from happening by verifying that I am who I say I am.

## Passwords

Passwords are a common way to authenticate users. You know how they work: When a user signs up for a new account, they choose a password. When they log in, they enter their password again. The server will then compare the password they entered with the password that was stored in the database.

There are 2 _really important_ things to consider when storing passwords:

1. **Storing passwords in plain text is awful.** If someone gets access to your database, they will be able to see all of your users' passwords. If you store passwords in plain text, you are giving away your users' passwords to anyone who gets access to your database.
2. **Password strength matters.** If you allow users to choose weak passwords, they will be more likely to reuse the same password on other websites. If someone gets access to your database, they will be able to log in to your users' other accounts.

We won't be writing code to validate password strength in this course, but you get the idea: you can enforce rules in your HTTP handlers to make sure passwords are of a certain length and complexity.

## Hashing

On the other hand, we _will_ be writing code to store passwords in a way that prevents them from being read by anyone who gets access to your database. This is called _hashing_. Hashing is a one-way function. It takes a string as input and produces a string as output. The output string is called a _hash_.

We'll cover how hashing works in-depth in a later course. For now, just know that hashing is a way to store passwords in a way that prevents them from being read by anyone who gets access to your database, but still allows us to _compare_ passwords when a user logs in.

## Assignment

1. [ ] Add and run a new migration that adds a non-null `TEXT` column to the `users` table called `hashed_password`. It should default to "unset" for existing users.

For the password hash, we will use [Argon2](https://en.wikipedia.org/wiki/Argon2). To help with this, use the library [argon2id](https://github.com/alexedwards/argon2id), which is a convenience wrapper around Argon2.

Download the library:

```bash
go get github.com/alexedwards/argon2id
```

2. [ ] Create an `internal/auth` package and expose two functions:
   - [ ] `func HashPassword(password string) (string, error)`: Hash the password using the `argon2id.CreateHash` function.
   - [ ] `func CheckPasswordHash(password, hash string) (bool, error)`: Use the `argon2id.ComparePasswordAndHash` function to compare the password that the user entered in the HTTP request with the password that is stored in the database.

I wrote a couple of simple [unit tests](https://go.dev/doc/tutorial/add-a-test) to ensure the package is working as expected.

3. [ ] Update the `POST /api/users` endpoint. The body parameters should now require a new `password` field:

```json
{
  "password": "04234",
  "email": "lane@example.com"
}
```

> [!info]
> As long as your server uses HTTPS in production, it's safe to send raw passwords in HTTP requests, because the entire request is encrypted.

Use your internal package's `HashPassword` function to hash the password before storing it in the database. Do **NOT** return the hashed password in the response. Again, that would be a security risk.

4. [ ] Add a `POST /api/login` endpoint. This endpoint should allow a user to login. In a future exercise, this endpoint will be used to give the user a token that they can use to make authenticated requests. For now, let's just make sure password validation is working. It should accept this body:

```json
{
  "password": "04234",
  "email": "lane@example.com"
}
```

You'll need a new query to look up a user by their email address (you don't have access to an ID here). Once you have the user, check to see if their password matches the stored hash using your internal package. If either the user lookup or the password comparison errors, just return a `401 Unauthorized` response with the message "Incorrect email or password".

If the passwords match, return a `200 OK` response and a copy of the user resource (without the password of course):

```json
{
  "id": "f0f87ec2-a8b5-48cc-b66a-a85ce7c7b862",
  "created_at": "2021-07-07T00:00:00Z",
  "updated_at": "2021-07-07T00:00:00Z",
  "email": "lane@example.com"
}
```

**Run and submit** the CLI tests.


## CLI

- Run: `bootdev run 294e5c16-d1e8-4836-871c-dedc98581236`
- Submit: `bootdev run 294e5c16-d1e8-4836-871c-dedc98581236 -s`
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
- Body contains: `id`
- Body contains: `created_at`
- Body contains: `updated_at`
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
- Body contains: `id`
- Body contains: `created_at`
- Body contains: `updated_at`
- JSON `.email` eq `saul@bettercall.com`

### Step 4
- **POST** `${baseURL}/api/login`
- Body:
```json
{
  "email": "saul@bettercall.com",
  "password": "000011112222"
}
```
- Expect status: 401

