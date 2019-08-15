package main

import (
	"fmt"
	"time"
)

func loadDinner(a appetizer, s salad, e entree) {
	result := fullDinner{
		name: fmt.Sprintf("%q with %q salad and a starter of %q.", e.name, s.name, a.name),
	}

	if a.hasCheese || e.hasCheese {
		result.hasCheese = true
	}

	if a.hasFruit || s.hasFruit {
		result.hasFruit = true
	}

	if s.isVegetarian && e.isVegetarian {
		result.isVegetarian = true
	}

	// Push to Remote data source
	time.Sleep(1250 * time.Millisecond)

	fmt.Println("Sent:", result.name)
}
