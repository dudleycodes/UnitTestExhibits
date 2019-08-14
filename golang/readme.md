# GO Testing Cheatsheet

## Unit Tests

| Action                       | Command                            | Notes                                                 |
| ---------------------------- | ---------------------------------- | ----------------------------------------------------- |
| Run tests in directory       | `go test .`                        |                                                       |
| Run tests in directory tree  | `go test ./...`                    |                                                       |
| Run tests matching regex     | `go test ./... --run=NameOfTest`   |                                                       |
| Run tests in specific file   | `go test ./file_test.go ./file.go` | Include the file the test depends on for compilation. |
| Clean test cache             | `go clean --testcache`             |                                                       |

## Code Coverage

> Remember: the only *absolute* conclusion that can be made is that code with zero coverage isn't unit tested at all. Code coverage does not guarantee the unit tests themselves are accurate and/or comprehensive.

| Action                        | Command                                     | Notes                                                                        |
| ----------------------------- | ------------------------------------------- | ---------------------------------------------------------------------------- |
| See basic summary             | `go test ./... --cover`                     |                                                                              |
| Save report to text file      | `go test ./... --coverprofile=coverage.out` | It's community convention to use the `.out` file extension for test results. |
| See summary of a saved report | `go tool cover --func=coverage.out`         |                                                                              |
| See details of a saved report | `go tool cover --html=coverage.out`         |                                                                              |

### Generate and View Code Coverage Heatmap

1. Generate and save coverage report with frequency information: `go test ./... --covermode=count --coverprofile=coverage.out`.
2. See details of the saved report: `go tool cover --html=coverage.out`.

## Benchmark Tests

> Beware high P-Values ([wikipedia](https://en.wikipedia.org/wiki/P-value)); benchmark results with a p-value greater than 0.05 are *probably* useless.

| Action                                       | Command                                            | Notes                                                                                                   |
| -------------------------------------------- | -------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| Run all benchmarks in directory              | `go test . --bench=.`                              | Also runs unit tests in directory.                                                                      |
| Run all benchmarks in directory tree         | `go test ./... --bench=.`                          | Also runs unit tests in directory tree.                                                                 |
| Run *only* benchmarks in directory tree      | `go test ./... --run=^$ --bench=. `                | Works by using a regex that matches nothing.                                                            |
| Run benchmarks matching regex                | `go test ./... --bench=NameOfTest`                 |                                                                                                         |
| Run benchmarks matching regex for 20 seconds | `go test ./... --bench=NameOfTest --benchtime=20s` | Increasing benchmark duration is *usually* the best way to reduce p-values for long-running functions.  |
| Run benchmarks matching regex 10 times       | `go test ./... --bench=NameOfTest --count=10`      | Increasing iterations is *usually* the best way to reduce p-values for short-running functions.         |

### Comparing Benchmarks

1. Install Russ Cox's [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) utility: `go get golang.org/x/perf/cmd/benchstat`.
2. Run the benchmark saving the results to a text file: `go test ./... --bench=NameOfTest --count=10 > original.out`.
3. Make your changes to the function under benchmark testing.
4. Re-run the *same* benchmark saving the results to a different text file: `go test ./... --bench=NameOfTest --count=10 > refactor.out`.
5. Compare using the benchstat util: `benchstat original.out refactor.out`.

## Additional Reading

- [Why bother writing tests at all?](https://dave.cheney.net/2019/05/14/why-bother-writing-tests-at-all) ([mirror](https://archive.is/h7nhI))

### Unit Testing

- [Golang basics - writing unit tests](https://blog.alexellis.io/golang-writing-unit-tests/) ([mirror](https://archive.is/AmW1I))
- [Prefer table driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests) ([mirror](https://archive.is/ufGU1))
- [Unit testing HTTP servers](https://www.youtube.com/watch?v=hVFEV-ieeew)
- [A pattern for Go tests](https://medium.com/@pierreprinetti/a-pattern-for-go-tests-3468b51535) ([mirror](https://archive.is/4aVtQ))

### Code Coverage

- [The Pitfalls of Code Coverage](https://blogs.msdn.microsoft.com/raulperez/2011/04/20/the-pitfalls-of-code-coverage/) ([mirror](https://archive.is/Q3xcO))

### Benchmark Testing

- [Analyzing the performance of Go functions with benchmarks](https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6) ([mirror](https://archive.is/VwUjW))
