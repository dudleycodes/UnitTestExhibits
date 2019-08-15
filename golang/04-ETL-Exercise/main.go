package main

import (
	"fmt"
)

func main() {
	// Appetizer
	fmt.Println("\nGetting appetizer from origin database....")
	a := extractAppetizer()
	fmt.Println("Got:", a.name)
	fmt.Println()

	// Salad
	fmt.Println("Getting salad from origin database...")
	s := extractSalad()
	fmt.Println("Got:", s.name)
	fmt.Println()

	// Entree
	fmt.Println("Getting entree from origin database...")
	e := extractEntree()
	fmt.Println("Got:", e.name)
	fmt.Println()

	// Loading
	fmt.Println("Sending courses to destination database...")
	loadDinner(a, s, e)
	fmt.Println()
}
