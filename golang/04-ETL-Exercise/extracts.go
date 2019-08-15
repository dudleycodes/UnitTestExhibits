package main

import "time"

func extractAppetizer() appetizer {
	// Mock function; normally would return from remote data source
	time.Sleep(1000 * time.Millisecond)

	return appetizer{
		hasCheese: true,
		hasFruit:  false,
		name:      "Cheesey Bread",
	}
}

func extractSalad() salad {
	// Mock function; normally would return from remote data source
	time.Sleep(750 * time.Millisecond)

	return salad{
		hasFruit:     false,
		isVegetarian: true,
		name:         "Summer Asian Slaw",
	}
}

func extractEntree() entree {
	// Mock function; normally would return from remote data source
	time.Sleep(1250 * time.Millisecond)

	return entree{
		hasCheese:    false,
		isVegetarian: false,
		name:         "Butter-Roasted Rib Eye Steak",
	}
}
