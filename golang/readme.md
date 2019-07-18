# GO(lang) Unit Test Examples: Basics

## Command Cheatsheet

### Run unit tests in current directory

```shell
go test .
```

### Clean test cache; run unit test in current directory

```shell
go clean -testcache && go test .
```

### Run benchmark tests in current directory

```shell
go test -bench=.
```

### Run unit tests for current directory and all sub-directories

```shell
go test ./...
```

## Additional Reading

- [Why bother writing tests at all?](https://dave.cheney.net/2019/05/14/why-bother-writing-tests-at-all)

### Unit Testing

- [Golang basics - writing unit tests](https://blog.alexellis.io/golang-writing-unit-tests/)
- [Prefer table driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
- [A pattern for Go tests](https://medium.com/@pierreprinetti/a-pattern-for-go-tests-3468b51535)

### Benchmark Testing

- [Analyzing the performance of Go functions with benchmarks](https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6)
