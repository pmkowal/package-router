package requests

import (
	"fmt"
	"packageRouter/internal/models"
	"testing"
)

func TestOSRMRequestModel_URL(t *testing.T) {
	osrmEndpoint := "https://router.project-osrm.org"
	testCases := []struct {
		Model OSRMRequestModel
		ExpectedResult string
	}{
		{Model: OSRMRequestModel{
			models.LocationModel{Latitude: 13.38886, Longitude: 52.517037},
			models.LocationModel{Latitude: 13.397634, Longitude: 52.529407},
		}, ExpectedResult: "https://router.project-osrm.org/route/v1/driving/13.388860,52.517037;13.397634,52.529407?overview=false"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("ShouldReturn`%s`From`%+v`", tc.ExpectedResult, tc.Model), func(t *testing.T) {
			result, _ := tc.Model.URL(osrmEndpoint)
			if result != tc.ExpectedResult {
				t.Fatalf("OSRMRequestModel.URL() returned\n`%s`\n, expected\n`%s`", result, tc.ExpectedResult)
			}
		})
	}
}
