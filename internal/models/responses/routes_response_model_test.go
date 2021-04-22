package responses

import (
	"reflect"
	"testing"
)

func TestRouteResponseModel_SortRoutesByDurationAndDistance(t *testing.T) {
	initialOrder := []RouteResponseModel{
		{"", 1444.5, 1700.1},
		{"", 251.5, 1884.8},
		{"", 384.5, 3795.2},
		{"", 384.5, 3695.2},
	}
	expectedOrder := []RouteResponseModel{
		{"", 251.5, 1884.8},
		{"", 384.5, 3695.2},
		{"", 384.5, 3795.2},
		{"", 1444.5, 1700.1},
	}
	responseModel := &RoutesResponseModel{}
	responseModel.Routes = initialOrder
	responseModel.SortRoutesByDurationAndDistance()
	if !reflect.DeepEqual(responseModel.Routes, expectedOrder) {
		t.Fatalf("RouteResponseModel.SortRoutesByDurationAndDistance() returned routes in a wrong order:\n`%+v`\n, expected\n`%+v`", initialOrder, expectedOrder)
	}
}