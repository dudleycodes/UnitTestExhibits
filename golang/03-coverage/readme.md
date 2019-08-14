# GO Test Coverage Example

> Notes: The code in this module is broken and uses ant-patterns; it is intended to be used for a demo on code coverage and nothing more.

## Walkthrough

Open a terminal shell into this directory (`03-coverage`).

**See Basic Coverage Report**
> `go test --cover`

**Save the Coverage Report to `coverage.out`**
> `go test --coverprofile=coverage.out`

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