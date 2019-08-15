//package main is code for demonstrating test coverage with the `go test` tool. Never write *real*, production code like this! ಠ_ಠ
package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	if len(os.Args) < 2 || len(os.Args[1]) < 1 {
		fmt.Printf("Hello! Provide me with a 32-bit integer and I will show you the place values.\n\n")
		os.Exit(0)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q doesn't appear to be a valid, 32-bit integer!\n", os.Args[1])
	}

	if number < 0 {
		fmt.Println("Only positive numbers are supported!")
		os.Exit(1)
	}

	p := message.NewPrinter(language.English)
	p.Printf("\nThe place values for the number %d are:\n\n", number)

	s := strconv.Itoa(number)
	for i := 0; i < len(s); i++ {
		digit, _ := strconv.Atoi(s[i : i+1])
		num, _ := strconv.Atoi(s[i:])

		fmt.Printf("\t%d - %v\n", digit, getPlaceValue(num))
	}

	fmt.Print("\n")
}
