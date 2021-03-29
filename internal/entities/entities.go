package entities

type Ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Plate struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	OnlyOn string `json:"only_on"`
}

type Step struct {
	Ingredient Ingredient `json:"ingredient"`
	Amount     float64    `json:"amount"`
	Unit       string     `json:"unit"`
}
