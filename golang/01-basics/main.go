package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	allowNegativeNumbers = true
	addDelay             = false
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Printf("Hello! Provide me with 32-bit integers and I will add them together for you.\n\n")
		os.Exit(0)
	}

	hasErr := false
	var integers []int
	for _, arg := range argsWithoutProg {
		if i, err := strconv.Atoi(arg); err != nil {
			hasErr = true
			fmt.Printf("%q doesn't appear to be a valid, 32-bit integer!\n", arg)
		} else {
			integers = append(integers, i)
		}
	}

	if hasErr {
		os.Exit(1)
	}

	sum := addNumbers(integers...)
	fmt.Printf("Those integers add up to: %d\n", sum)
}

func addNumbers(nums ...int) int {
	r := 0

	for _, n := range nums {
		//For demonstrating test results
		if !allowNegativeNumbers && n < 0 {
			return 0
		}

		r += n
	}

	//For demonstrating benchmarking results
	if addDelay {
		time.Sleep(20 * time.Millisecond)
	}

	return r
}
