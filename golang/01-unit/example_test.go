package main_test

import (
	"fmt"

	"github.com/dudleycodes/UnitTestExhibits/01-unit/main"
)

func Example_Sum() {
	sum := main.Sum(2, 4, 6, 8, 12)
	fmt.Println(sum)
	// Output: 32
}
