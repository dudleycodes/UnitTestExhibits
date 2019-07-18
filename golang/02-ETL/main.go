package main

import (
	"fmt"
)

func main() {
	// Entree
	fmt.Println("Getting entree from origin database....")
	e := extractEntree()
	fmt.Println("Got:", e.name)
	fmt.Println()

	// Salad
	fmt.Println("Getting salad from origin database...")
	s := extractSalad()
	fmt.Println("Got:", s.name)
	fmt.Println()

	// Main Dish
	fmt.Println("Getting main dish from origin database...")
	d := extractMainDish()
	fmt.Println("Got:", d.name)
	fmt.Println()

	// Loading
	fmt.Println("Sending courses to destination database...")
	loadDinner(e, s, d)
	fmt.Println()
}
