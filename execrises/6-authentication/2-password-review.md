# Password Review

_Chapter: Authentication | Slug: 2-password-review | UUID: 39b70a72-d801-4f8f-bccf-97a4064b62be_

# Password Review

It's a really bad idea for users to reuse the same passwords across sites. If someone figures out their password for one site, they can try it on other sites. If they get lucky, they can log in to and compromise many of their accounts.

Unfortunately, it's very common for users to reuse passwords. We can't _force_ users to not reuse passwords on the server side, but we can take steps to make it harder for them to reuse passwords. Namely, we can require that passwords are strong.

## Passwords Should Be Strong

The most important factor for the strength of a password is its _entropy_. [Entropy](https://www.boot.dev/blog/computer-science/what-is-entropy-in-cryptography/) is a measure of how many possible combinations of characters there are in a string. To put it simply:

- The longer the password the better
- Special characters and capitals should always be allowed
- Special characters and capitals aren't as important as length

![password strength](https://imgs.xkcd.com/comics/password_strength.png)

- [xkcd: Password Strength](https://xkcd.com/936/)

## Passwords Should Never Be Stored in Plain Text

The most critical thing we can do to protect our users' passwords is to _never_ store them in plain text. We should use cryptographically strong key derivation functions (which are a special class of hash functions) to store passwords in a way that prevents them from being read by anyone who gets access to your database.

[Argon2id](https://en.wikipedia.org/wiki/Argon2) is a great choice. [SHA-256](https://www.boot.dev/blog/computer-science/how-sha-2-works-step-by-step-sha-256/) and [MD5](https://en.wikipedia.org/wiki/MD5) are not.

## Question

**Which is a good hash function for passwords?**

- SHA-256
- MD5
- Argon2id ✅

**Answer:** Argon2id
