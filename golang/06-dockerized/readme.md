# Go - Automated Integration Test with Docker Containers

## About

This example demonstrates automated, integration tests that involves spinning up and testing against
a Docker container. For more information see the blog post [Automating Go Integration Tests With
Docker](https://www.dudley.codes/posts/2020.10.02-golang-docker-integration-tests/).

## Walkthrough

To run this example Docker must be installed.

```shell
go test ./... --tags integration --count=1
```
