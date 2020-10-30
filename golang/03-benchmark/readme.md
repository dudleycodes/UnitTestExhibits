# GO Benchmark Test Examples

## Walkthrough

### Run the stock benchmark tests

```shell
go test --bench=.
```

### Optimize and Re-run

Comment out the `time.Sleep` statement in the bottom `fib()` function and rerun the benchmark test:
`go test --bench=.`  You will see the performance improve.

### Let's increase the accuracy

The computer has lots of stuff going on at any given moment, so let's increase the accuracy of our
results by running the test for multiple iterations.

`go test --bench=. --count=6`

### Save the Results

Let's re-run the benchmark and save the results to a file so that we can compare future results:
`go test --bench=. --count=6 | tee original.txt`.

### Refactor for speed and Re-test

Using block comments (`/*` and `*/`), comment out the bottom `fib()` function and un-comment the top
`fib()` function.  Re-run the test, saving the results to a new results file:

```shell
go test --bench=. --count=6 | tee refactor.txt
```

### Compare the Results

If you haven't already, install the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat)
utility via `go get golang.org/x/perf/cmd/benchstat`.

Now let's compare the results of our two tests: `benchstat original.txt refactor.txt`

## Warnings

**THE BEST OPTIMIZATIONS ARE THOSE THAT MAKE CODE THE EASIEST TO MAINTAIN AND THE EASIEST TO
REFACTOR/DELETE**  See also: [Write code that is easy to delete, not easy to extend.](https://programmingisterrible.com/post/139222674273/write-code-that-is-easy-to-delete-not-easy-to) ([mirror](https://archive.fo/dJqT4)).

Optimizing code too early, particularly code that's likely to change, is a terrible mistake:

> "We should forget about small efficiencies, say about 97% of the time: **premature optimization is
> the root of all evil**. Yet we should not pass up our opportunities in that critical 3%" -- Donald
> Knuth

Faster code is usually longer, harder to maintain, and difficult to debug:

> "Everyone knows that debugging is twice as hard as writing a program in the first place. So if
> you're as clever as you can be when you write it, how will you ever debug it?" --Brian Kernighan

Always optimize for our teammates first:

![Code Tip](https://raw.githubusercontent.com/dudleycodes/UnitTestExhibits/master/.misc/maniacCodeTip.jpg "Code Tip")

## Optimization Notes

Go does a lot of optimizations for us by default, and has many more optimizations that can be
toggled on. Understanding these optimizations is the best first step (see: [Go Compiler And Runtime
Optimizations](https://github.com/golang/go/wiki/CompilerOptimizations)).

- The ultimate analysis method is examining the generated GO code: `go tool compile -S fibonacci.go`
- Example of boundary checking
- Play with "function inlining"
- Reduce allocations and unnecessary casting

## See Also

These examples are heavily influenced/copied from Dave Cheney's ["High Performance GO
Workshop"](https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html#benchmarking),
licensed under the [Creative Commons Attribution-ShareAlike 4.0
International](https://creativecommons.org/licenses/by-sa/4.0/) licence.
