package models

type RoutesResponseModel struct {
	SourceRaw string `json:"source"`
	Routes []struct {
		DestinationRaw string  `json:"destination"`
		Duration    float64 `json:"duration"`
		Distance    float64 `json:"distance"`
	} `json:"routes"`
}