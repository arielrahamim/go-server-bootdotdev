# Types of Authentication

_Chapter: Authentication | Slug: 5--auth-types | UUID: eeb8e455-08f8-4711-a2d7-72089ae5dd66_

# Types of Authentication

Here are a few of the most common authentication methods you'll see in the wild:

1. Password + ID (username, email, etc.)
2. 3rd Party Authentication ("Sign in with Google", "Sign in with GitHub", etc)
3. Magic Links
4. API Keys

## 1. Password + ID

This is the most common type of authentication that requires a manual login from a user. When users use password managers, it's one of the more secure ways to authenticate users, unfortunately, many users don't, so it's not as secure as it could be.

That said, it's a valid choice.

## 2. 3rd Party Authentication

3rd party authentication is a way to authenticate users using a service like Google or GitHub. 3rd party auth is great for user experience because it allows users to use their existing accounts to log in to your app, lowering friction.

It's also nice because you don't need to worry about storing passwords yourself, meaning you can outsource the security of your users' passwords to a company that, _hopefully_, does a good job.

The only real drawbacks to 3rd party auth is that you're trusting a 3rd party and if your users don't have an account with that 3rd party, they won't be able to log in.

## 3. Magic Links

Magic links are a way to authenticate users without a password. It relies on the assumption that the user's email is something that they have unique access to.

The webserver sends a link to the user's email and encodes a unique token in that link. When the user clicks the link, the webserver can decode the token and use it to authenticate the user. Eg:

`https://example.com/login?token=...`

## 4. API Keys

API keys are a fantastic way to authenticate users and systems programmatically. An API Key is just a long, secure string that uniquely identifies a user or system, and that can't be guessed. Because they're intended to be used in code, they don't need to be memorized and, as such, can be much longer and double as an identifier. An API key might look something like this:

`bd_JDS543J3n5NMKspDXNRlowiqw523lKHK32K43kl`

![types of authentication](https://storage.googleapis.com/qvault-webapp-dynamic-assets/lesson_videos/types-of-authentication-v2-1920x1080.mp4)

## Question

**Which makes the most sense to authenticate a user of a command line application for secure API access?**

- Password + ID
- Magic link
- API key ✅

**Answer:** API key
