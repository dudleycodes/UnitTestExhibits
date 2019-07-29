package main

type appetizer struct {
	hasCheese bool
	hasFruit  bool
	name      string
}

type salad struct {
	hasFruit     bool
	isVegetarian bool
	name         string
}

type entree struct {
	hasCheese    bool
	isVegetarian bool
	name         string
}

type fullDinner struct {
	hasCheese    bool
	hasFruit     bool
	isVegetarian bool
	name         string
}
