# Cookies

_Chapter: Authentication | Slug: 13-cookies | UUID: 58d2fb99-b84f-4867-89e9-f9b46b51e93e_

# Cookies

HTTP [cookies](https://en.wikipedia.org/wiki/HTTP_cookie) are one of the most talked about, but least understood, aspects of the web.

When cookies are talked about in the news, they're usually implied to simply be privacy-stealing bad guys. While cookies can certainly invade your privacy, that's not what they _are_.

## What Is an HTTP Cookie?

A cookie is a small piece of data that a server sends to a client. The client then dutifully stores the cookie and sends it back to the server on subsequent requests.

Cookies can store any arbitrary data:

- A user's name or other tracking information
- A JWT (refresh and access tokens)
- Items in a shopping cart
- etc.

The server decides _what_ to put in a cookie, and the client's job is simply to store it and send it back.

## How Do Cookies Work?

Simply put, cookies work through HTTP headers.

Cookies are sent from the server to the client in the `Set-Cookie` header. Cookies are most popular for web (browser-based) applications because browsers _automatically_ send any cookies they have back to the server in the `Cookie` header.

## Why Aren't We Using Cookies?

Simply put, Chirpy's API is designed to be consumed by mobile apps and other servers. Cookies are primarily for browsers.

A good use-case for cookies is to serve as a more strict and secure transport layer for JWTs within the context of a browser-based application.

For example, when using [httpOnly cookies](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies#block_access_to_your_cookies), you can ensure that 3rd party JavaScript that's being executed on your website can't access any cookies. That's a lot better than storing JWTs in the browser's local storage, where it's easily accessible to any JavaScript running on the page.

## Question

**Cookies can store...**

- User preferences
- Session information
- Authentication information
- Anything ✅

**Answer:** Anything
