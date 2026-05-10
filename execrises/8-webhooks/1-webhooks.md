# Webhooks

_Chapter: Webhooks | Slug: 1-webhooks | UUID: 1304e939-bf50-48d3-a351-b35faafc267d_

# Webhooks

Webhooks sound like a scary advanced concept, but they're quite simple.

A webhook is just an event that's sent to your server by an external service when something happens.

For example, here at Boot.dev we use Stripe as a third-party payment processor. When a student makes a payment, Stripe sends a webhook to the Boot.dev servers so that we can unlock the student's membership.

1. Student makes a payment to stripe
2. Stripe processes the payment
3. If the payment is successful, Stripe sends an `HTTP POST` request to `https://api.boot.dev/stripe/webhook` (that's not the real URL, but you get the idea)

That's it! The only real difference between a webhook and a typical `HTTP` request is that the system making the request is an automated system, not a human loading a webpage or web app. As such, webhook handlers must be [idempotent](https://en.wikipedia.org/wiki/Idempotence) because the system on the other side may retry the request multiple times.

## Idempo... What?

Idempotent, or "idempotence", is a fancy word that means "the same result no matter how many times you do it". For example, your typical `POST /api/chirps` (create a chirp) endpoint will _not_ be idempotent. If you send the same request twice, you'll end up with two chirps with the same information but different IDs.

Webhooks, on the other hand, should be idempotent, and it's typically easy to build them this way because the client sends some kind of "event" and usually provides its own unique ID.

## Assignment

We recently rolled out a new feature called "Chirpy Red". It's a membership program, and members of "Chirpy Red" get pretty incredible features: like the ability to edit chirps after posting them. But that's beside the point...

Chirpy uses "Polka" as its payment provider. They send us webhooks whenever a user subscribes to Chirpy Red. We need to mark users as Chirpy Red members when we receive these webhooks.

1. [ ] Add a migration to the `users` table to include a new column called `is_chirpy_red`. This column should be a boolean, and it should default to `false`.
2. [ ] Add a database query that upgrades a user to chirpy red based on their ID.
3. [ ] Add a `POST /api/polka/webhooks` endpoint. It should accept a request of this shape:

```json
{
  "event": "user.upgraded",
  "data": {
    "user_id": "3311741c-680c-4546-99f3-fc9efac2036c"
  }
}
```

- [ ] If the `event` is anything _other_ than `user.upgraded`, the endpoint should immediately respond with a `204` status code - we don't care about any other events.
- [ ] If the `event` _is_ `user.upgraded`, then it should update the user in the database, and mark that they are a Chirpy Red member.
- [ ] If the user is upgraded successfully, the endpoint should respond with a `204` status code and an empty response body. If the user can't be found, the endpoint should respond with a `404` status code.

_Polka uses the response code to know whether or not the webhook was received successfully. If the response code is anything other than `2XX`, they'll retry the request._

4. [ ] Update all endpoints that return user resources to include the `is_chirpy_red` field.

**Run and submit** the CLI tests.


## CLI

- Run: `bootdev run 1304e939-bf50-48d3-a351-b35faafc267d`
- Submit: `bootdev run 1304e939-bf50-48d3-a351-b35faafc267d -s`
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
- JSON `.is_chirpy_red` eq `None`
- Capture variable `userID` from `.id`

### Step 3
- **POST** `${baseURL}/api/polka/webhooks`
- Body:
```json
{
  "data": {
    "user_id": "${userID}"
  },
  "event": "user.payment_failed"
}
```
- Expect status: 204

### Step 4
- **POST** `${baseURL}/api/polka/webhooks`
- Body:
```json
{
  "data": {
    "user_id": "${userID}"
  },
  "event": "user.upgraded"
}
```
- Expect status: 204

### Step 5
- **POST** `${baseURL}/api/login`
- Body:
```json
{
  "email": "walt@breakingbad.com",
  "password": "123456"
}
```
- Expect status: 200
- JSON `.is_chirpy_red` eq `None`

