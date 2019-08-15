# GO Test Coverage Example

## About

This example demonstrates generate and reading code coverage reports.  The code in this module is *purposely* broken and uses anti-patterns -- it should not be used for anything real.

## Walkthrough

Open a terminal shell into this directory (`03-coverage`).

**See Basic Coverage Report**
> `go test . --cover`

**Save the Coverage Report to `coverage.out`**
> `go test . --coverprofile=coverage.out`

**See the Basic `coverage.out` Contents**
> `go tool cover --func=coverage.out`

**See Report on `coverage.out` Contents**
> `go tool cover --html=coverage.out`

**Generate and See Heat-map Coverage Results**
> ```shell
> go test --covermode=count --coverprofile=coverage.out
> go tool cover --html=coverage.out
> ```

## Further Reading
- [The cover story](https://blog.golang.org/cover)
