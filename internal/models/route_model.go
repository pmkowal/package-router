package models

type RouteModel struct {
	DestinationRaw string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}