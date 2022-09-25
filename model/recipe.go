package model

type Recipe struct {
	name            string
	healthScore     int
	aggragatedLikes int
	servings        int
	glutenFree      bool
	pricePerServing float64
	readyInMinutes  int
}
