package models

import (
	"reflect"
	"testing"
)

func TestSortRoutesByDurationAndDistance(t *testing.T) {
	routes := []RouteResponseModel {
		{"", 1444.5, 1700.1},
		{"", 251.5, 1884.8},
		{"", 384.5, 3795.2},
		{"", 384.5, 3695.2},
	}
	sortedRoutes := []RouteResponseModel {
		{"", 251.5, 1884.8},
		{"", 384.5, 3695.2},
		{"", 384.5, 3795.2},
		{"", 1444.5, 1700.1},
	}
	responseModel := &RoutesResponseModel{}
	responseModel.Routes = routes
	responseModel.SortRoutesByDurationAndDistance()
	if !reflect.DeepEqual(responseModel.Routes, sortedRoutes) {
		t.Fatalf("SortRoutesByDurationAndDistance() returned routes in a wrong order:\n`%+v`\n, expected\n`%+v`", routes, sortedRoutes)
	}
}