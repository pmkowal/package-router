package responses

import (
	"math/big"
	"packageRouter/internal/models"
	"sort"
)

type RoutesResponseModel struct {
	SourceRaw string               `json:"source"`
	Routes    []RouteResponseModel `json:"routes"`
}

func (m *RoutesResponseModel) SortRoutesByDurationAndDistance() {
	sort.Slice(m.Routes, func(i, j int) bool {
		leftDuration := m.Routes[i].Duration
		rightDuration := m.Routes[j].Duration
		switch big.NewFloat(leftDuration).Cmp(big.NewFloat(rightDuration)) {
		case -1:
			return true
		case 1:
			return false
		}
		leftDistance := m.Routes[i].Distance
		rightDistance := m.Routes[j].Distance
		switch big.NewFloat(leftDistance).Cmp(big.NewFloat(rightDistance)) {
		case -1:
			return true
		}
		return false
	})
}

type RouteResponseModel struct {
	DestinationRaw string  `json:"destination"`
	Duration       float64 `json:"duration"`
	Distance       float64 `json:"distance"`
}

func (m *RouteResponseModel) SetRoute(routeModel *models.RouteModel) {
	m.Duration = routeModel.Duration
	m.Distance = routeModel.Distance
}
