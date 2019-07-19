package main

import (
	"fmt"
	"time"
)

func loadDinner(e entree, s salad, d mainDish) {
	result := fullDinner{
		name: fmt.Sprintf("%q with %q salad and a starter of %q.", d.name, s.name, e.name),
	}

	if e.hasCheese || d.hasCheese {
		result.hasCheese = true
	}

	if e.hasFruit || s.hasFruit {
		result.hasFruit = true
	}

	if s.isVegetarian && d.isVegetarian {
		result.hasFruit = true
	}

	// Push to Remote data source
	time.Sleep(1250 * time.Millisecond)

	fmt.Println("Sent:", result.name)
}
