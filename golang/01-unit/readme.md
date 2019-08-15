# GO Unit Test Example

## About

This example shows some different methods of writing unit test for a simple function; along with their pros and cons. The app under testing takes a series of 32-bit integers from the command line, and displays their sum and average.

```shell
MASCHEEN:01-unit user$ go run . 8 6 7 5 3 0 9
The average of those integers is: 5
Those integers add up to: 38
```

## Walkthrough

### Test 1 - Very Basic

This test does its job but when it fails we don't get any information about why or how.

```shell
go test . --run=01
```

### Test 2 - Tedious

This test works and give us some information about what failed but the code is tedious to not only to add/edit/remove test cases, but also to read.

```shell
go test . --run=02
```

### Test 3 - Less Tedious

This test works and cuts out a lot of the tedium but:
- it's hard to add/edit/remove test cases because it mixes the test-data with logic. When writing tests we only want to focus on the testing data (inputs/outputs).
- the test results still don't show us what inputs break things (only shows expected/actual)

```shell
go test . --run=03
```

### Test 4 - Table Driven Test

This "table test" works; by isolating the logic we can focus just on test data when add/editing/removing. More importantly we can show all data round test failures.

```shell
go test . --run=04
```

### Test 5 - Named Table Driven Test

Named table tests have all the benefits of normal table tests. By providing each test case with a name we can ensure the motivations behind the tests are also displayed in the test results (and not just code comments like the earlier examples).

```shell
go test . --run=05
```

## Exercise

1. Add a table-driven test for the `average` function; make sure the tests trigger each line of code.

2. Refactor the code that validates and converts command line arguments into its own function and add to unit testing. Make sure to have the following test cases:

   - Accepts a single integer
   - Accepts multiple integers
   - Rejects all non-integer types (string, int64, float, etc)
