# Webhooks Review

_Chapter: Webhooks | Slug: 3--webhooks-review | UUID: f7651474-4443-41ba-a74e-aeccc1a17372_

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

**What will likely happen if your server responds to a webhook with a 5XX status code?**

- The third party will stop trying to send the event
- The third party will never send another event to your server
- The third party will keep retrying the exact same request ✅

**Answer:** The third party will keep retrying the exact same request
