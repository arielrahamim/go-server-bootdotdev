# Documentation

_Chapter: Documentation | Slug: 3--documentation | UUID: 73d5365b-467e-42db-9ca4-f34a246aa82e_

# Documentation

As far as creating documentation goes, there are 2 main approaches:

1. Manually write documentation
2. Use a tool to generate documentation

Obviously, the first approach is easier to get going with if you have a small API, but as the system grows, it can be really hard to keep the documentation up to date.

**Incorrect documentation is worse than no documentation.**

At least when there is _no_ documentation, your clients will reach out and ask for clarification. When the documentation is incorrect, it can lead to a lot of wasted time and frustration.

## Manually Writing Documentation

When I've worked on smaller teams, we've generally opted to write our documentation in [Markdown files](https://www.markdownguide.org/) and host them on GitHub. This is a great way to get started because Markdown is a simple format that is easy to write and easy to read.

## Automated Documentation Generation

I've also written and consumed APIs that have used:

- [Swagger](https://swagger.io/)
- [GraphQL](https://graphql.org/) (not RESTful, but still a networking API)
- [Godoc](https://go.dev/blog/godoc) (which only works for REST APIs if you provide an SDK)
- [Postman](https://learning.postman.com/docs/publishing-your-api/documenting-your-api/) (only useful if your team all uses Postman as their HTTP client)

And with LLMs, there are also now trivial ways to use AI to generate simple markdown documentation directly from your backend code.

## Okay, but What Should I Do for Now?

I recommend personally writing documentation for your projects in Markdown files and storing them alongside the rest of your code in Git. Your project's `README.md` file is a great place to start, but it's also common for the `README.md` file to link to a `/docs` folder that contains more detailed documentation. The benefits are:

- It's easy to get started (no need to install any additional tools)
- The documentation lives alongside your code, so it's easy to keep it up to date
- You'll learn Markdown, which is a great skill to have
- GitHub/GitLab will render your Markdown files for you, so your docs will look great

## Question

**Which is easier to keep up-to-date?**

- Documentation that lives in the same repo as the code ✅
- Documentation that lives in a separate repo/site/wiki

**Answer:** Documentation that lives in the same repo as the code
