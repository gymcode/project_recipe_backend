package model

type Recipe struct {
	ID              int
	Name            string
	HealthScore     int
	AggragatedLikes int
	Servings        int
	GlutenFree      bool
	PricePerServing float64
	ReadyInMinutes  int
	Summary         string
	Title           string
	Vegetarian      bool
	Vegan           bool
	VeryHealth      bool
	Verypopular     bool
	Cheap           bool
	CreatedAt       string
	UpdatedAt       string
}
