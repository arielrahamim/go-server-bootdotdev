Boot.dev
Dashboard
Courses
Training
Billing
Community
Leaderboard

New NotificationsToggle notifications6
gem bag
Disciple

Level 41

user avatarprofile role frame

sharpshooter armor
sharpshooter
16

streak embers

daily streak
Explain difficulty
xp potions

chest!Quest available

CH5: Storage

L9: Get All Chirps
Back
Next

Get All Chirps
We need a way to retrieve all the chirps from the database. Later, we'll add sorting and filtering functionality, but you can think of this as a very basic version of an endpoint that might serve a timeline of chirps.

Assignment




Boots
Spellbook
Lessons
Boots
Need help? I, Boots the Primeval 10x Developer, can assist... for a price.

Ask Boots a question...

Copy/paste one of the following commands into your terminal:

Run

bootdev run 341b80d4-556f-4c5b-8afc-ffd12d5238c2


Submit

bootdev run 341b80d4-556f-4c5b-8afc-ffd12d5238c2 -s


Default Base URL: http://localhost:8080
Optionally configure your CLI to override the default base URL by running bootdev config base_url <url>
Run the CLI commands to test your solution.

POST /admin/reset
1. Expecting status code: 200
POST /api/users
Request Body:
{
  "email": "saul@bettercall.com"
}

1. Expecting status code: 201
Parsing userID1 variable from response body .id
POST /api/chirps
Request Body:
{
  "body": "If you're committed enough, you can make any story work.",
  "user_id": "${userID1}"
}

1. Expecting status code: 201
POST /api/chirps
Request Body:
{
  "body": "I once told a woman I was Kevin Costner, and it worked because I believed it.",
  "user_id": "${userID1}"
}

1. Expecting status code: 201
GET /api/chirps
1. Expecting status code: 200
2. Expecting JSON at .[0].body to be equal to If you're committed enough, you can make any story work.
3. Expecting JSON at .[1].body to be equal to I once told a woman I was Kevin Costner, and it worked because I believed it.

Solution Files
Using the Bootdev CLI

The Bootdev CLI is the only way to submit your solution for this type of lesson. We need to be able to run commands in your environment to verify your solution.

You can install it here. It's a Go program hosted on GitHub, so you'll need Go installed as well. Instructions are on the GitHub page.

Oops, lost connection! Click here to refresh!
×