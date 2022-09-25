package model

type Instruction struct {
	Id        int
	Name      string
	Steps     []Steps
	CreatedAt string
	UpdatedAt string
}

type Steps struct {
	Id         int
	Number     int
	Equipment  []Equipments
	Step       string
	Ingredient []Ingredients
	CreatedAt  string
	UpdatedAt  string
}

type Equipments struct {
	Id            int
	Name          string
	EquipmentType string
	CreatedAt     string
	UpdatedAt     string
}
