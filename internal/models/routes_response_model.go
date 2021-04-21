package models

type RoutesResponseModel struct {
	SourceRaw string               `json:"source"`
	Routes    []RouteResponseModel `json:"routes"`
}

type RouteResponseModel struct {
	DestinationRaw string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

func (m *RouteResponseModel) SetRoute(routeModel *RouteModel) {
	m.Duration = routeModel.Duration
	m.Distance = routeModel.Distance
}
