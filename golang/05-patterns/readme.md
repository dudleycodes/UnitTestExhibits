# Useful Unit Test Patterns for Go

## Mutator-Table Driven Test

[`mutator.go`](`mutator.go`)

A derivative of table-driven tests, which contains a table of mutators, each which mutates a single
field of a struct. Useful for targeting functions that validate structs.
