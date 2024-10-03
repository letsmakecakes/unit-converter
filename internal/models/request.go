package models

type Request struct {
	FromUnit string  `json:"from_unit"`
	ToUnit   string  `json:"to_unit"`
	Value    float64 `json:"value"`
}
