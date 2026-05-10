# JWT Review

_Chapter: Authentication | Slug: 8-jwt-review | UUID: 4ba1bb2e-4f8a-46a0-89db-66f5181e8441_

# JWT Review

JWTs are cryptographically signed JSON objects that contain information about an authenticated user.

> [!lane]
> I've heard "JWT" pronounced as "jot", but I pronounce it "jay double yoo tee".

## JWTs Can't Be Changed

We'll talk about [MACs, HMACs](https://www.boot.dev/blog/backend/hmac-and-macs-in-jwts/), and digital signatures in a later course, which are the cryptographic concepts that power JWTs. For now, it's just important to know that once the token is created by a server, the data in the token can't be changed without the server being aware of it.

_When your server issues a JWT to Bob, Bob can use that token to make requests as Bob to your API. Bob won't be able to change the token to make requests as Alice._

## JWTs Are Not Encrypted

JWTs are not encrypted. Anyone who has the token can read the data (like the expiry and the user id) in the token. This is why you should never store sensitive information in a JWT. It's just a way to authenticate a user.

I like using [JWT.io](https://jwt.io/) to inspect JWTs. It is a great tool for playing around with them and learning how they work.

## JWT Lifecycle

![jwt lifecycle](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/hFgop3U-480x720.png)

## Question

**The data stored in a JWT is encrypted so that only the server can read it.**

- True
- False ✅

**Answer:** False
