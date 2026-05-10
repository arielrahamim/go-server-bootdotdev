# Webhooks Review

_Chapter: Webhooks | Slug: 2-webhooks-review | UUID: 581af9ab-b4a1-4477-b3e3-fe13ef8645e7_

# Webhooks Review

A webhook is just an event that's sent to your server by an external service. There are just a couple of things to keep in mind when building a webhook handler:

- The third-party system will probably retry requests multiple times, so your handler should be [idempotent](https://en.wikipedia.org/wiki/Idempotence).
- Be extra careful to never "acknowledge" a webhook request unless you processed it successfully. By sending a `2XX` code, you're telling the third-party system that you processed the request successfully, and they'll stop retrying it.
- When you're writing a server, you typically get to define the API. However, when you're integrating a webhook from a service like Stripe, you'll probably need to adhere to their API: they'll tell you what shape the events will be sent in.

## Are Webhooks and Websockets the Same Thing?

Nope! A websocket is a persistent connection between a client and a server. Websockets are typically used for real-time communication, like chat apps. Webhooks are a one-way communication from a third-party service to your server.

We'll talk about websockets in a future course.

![whats a webhook](https://storage.googleapis.com/qvault-webapp-dynamic-assets/lesson_videos/webhooks-1920x1080.mp4)

## Question

**Webhooks are `____`, websockets are `____`**

- Standard HTTP requests, persistent connections ✅
- Persistent connections, standard HTTP requests

**Answer:** Standard HTTP requests, persistent connections
