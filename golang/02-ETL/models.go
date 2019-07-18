package main

type entree struct {
	hasCheese bool
	hasFruit  bool
	name      string
}

type salad struct {
	hasFruit     bool
	isVegetarian bool
	name         string
}

type mainDish struct {
	hasCheese    bool
	isVegetarian bool
	name         string
}

type fullDinner struct {
	hasCheese    bool
	hasFruit     bool
	isVegatarian bool
	name         string
}
