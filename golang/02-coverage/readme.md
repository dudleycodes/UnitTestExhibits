# GO Test Coverage Example

## About

This example demonstrates generate and reading code coverage reports. The app under testing takes a 32-bit integer from the command line and displays the place values for each digit.

```shell
MASCHEEN:01-unit user$ go run . 1984

The place values for the number 1,984 are:

        1 - thousands
        9 - hundreds
        8 - tens
        4 - ones
```

***The code in this module is *intentionally* broken and uses anti-patterns -- it should not be used for anything real.***

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
