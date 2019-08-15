package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Printf("Hello! Provide me with 32-bit integers and I will add them together for you.\n\n")
		os.Exit(0)
	}

	var integers []int
	for _, arg := range argsWithoutProg {
		i, err := strconv.Atoi(arg)

		if err != nil {
			fmt.Printf("%q is not a valid 32-bit integer!\n", arg)
			os.Exit(1)
		}

		integers = append(integers, i)
	}

	avg := Average(integers...)
	fmt.Printf("The average of those integers is: %d\n", avg)

	sum := Sum(integers...)
	fmt.Printf("Those integers add up to: %d\n", sum)
}

// Average calculates the average value of a series of 32-bit integers.
func Average(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := Sum(nums...)

	return sum / len(nums)
}

// Sum calculates the sum value of a series of 32-bit integers.
func Sum(nums ...int) int {
	r := 0

	for _, n := range nums {
		//For demonstrating test results
		if n < 0 {
			return 0
		}

		r += n
	}

	return r
}
