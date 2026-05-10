# JWTs

_Chapter: Authentication | Slug: 6-jwt | UUID: be93db0d-4c6d-49cf-b56d-ba22392eb160_

# JWTs

There are several different ways to handle authentication. We'll use [JWTs](https://www.boot.dev/blog/backend/hmac-and-macs-in-jwts/) in this course. They're a popular choice for APIs that are consumed by web applications and mobile apps.

## What Is a JWT?

A JWT is a JSON Web Token. It's a cryptographically signed JSON object that contains information about the user. You'll learn about how the cryptography of JWTs work in our [Learn Cryptography](https://boot.dev/courses/learn-cryptography) course, for now, it's just important to know that once the token is created by the server, the data in the token can't be changed without the server knowing.

_When your server issues a JWT to Bob, Bob can use that token to make requests as Bob to your API. Bob won't be able to change the token to make requests as Alice._

## Assignment

The first building blocks you'll write are the functions for creating and validating JWTs, which will be used in the next lesson to authenticate users.

1. [ ] Add a `MakeJWT` function to your `auth` package:

```go
func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error)
```

Create and return a JWT using this [JWT library](https://github.com/golang-jwt/jwt), which you can import into your code by running:

```sh
go get -u github.com/golang-jwt/jwt/v5
```

2. [ ] Create a new token.
   1. [ ] Use [`jwt.NewWithClaims`](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#NewWithClaims)
   2. [ ] Use [`jwt.SigningMethodHS256`](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#SigningMethodHS256) as the signing method.
   3. [ ] Use [`jwt.RegisteredClaims`](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#RegisteredClaims) as the claims.
      - [ ] Set the `Issuer` to "chirpy-access"
      - [ ] Set `IssuedAt` to the current time in UTC
      - [ ] Set `ExpiresAt` to the current time plus the expiration time (`expiresIn`)
      - [ ] Set the `Subject` to a stringified version of the user's `id`
   4. [ ] Use [`token.SignedString`](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Token.SignedString) to sign the token with the secret key. Refer to [here](https://golang-jwt.github.io/jwt/usage/signing_methods/#signing-methods-and-key-types) for an overview of the different signing methods and their respective key types.
3. [ ] Add a `ValidateJWT` function to your `auth` package:

```go
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error)
```

4. [ ] Use the [`jwt.ParseWithClaims`](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#ParseWithClaims) function to validate the signature of the JWT and extract the claims into a [`*jwt.Token`](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Token) struct. The `keyFunc` callback must return the same key type (`[]byte`) used when the token was signed. An error will be returned if the token is invalid or has expired.

If all is well with the token, use the [`token.Claims`](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Claims) interface to get access to the user's `id` from the claims (which should be stored in the `Subject` field). Return the `id` as a `uuid.UUID`.

5. [ ] Add some more unit tests to the `auth` package. Make sure that you can create and validate JWTs, and that expired tokens are rejected and JWTs signed with the wrong secret are rejected.

**Run and submit** the CLI tests.

## Test Steps

### Step 1
```
go test ./...
```

