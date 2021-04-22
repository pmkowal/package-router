package models

import (
	"fmt"
	"testing"
)

func TestLocationModel_Parse(t *testing.T) {
	testCases := []struct {
		Location       string
		ExpectedResult LocationModel
	}{
		{"13.388860,52.517037", LocationModel{13.388860, 52.517037}},
		{"13.388860", LocationModel{}},
		{"m13.388860,52.517037", LocationModel{}},
		{"malformedString", LocationModel{}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("ShouldReturn`%+v`From`%s`", tc.ExpectedResult, tc.Location), func(t *testing.T) {
			locationModel := LocationModel{}
			locationModel.Parse(tc.Location)
			if locationModel != tc.ExpectedResult {
				t.Fatalf("LocationModel.Parse() returned `%+v` , expected `%+v`", locationModel, tc.ExpectedResult)
			}
		})
	}
}

func TestLocationModel_Description(t *testing.T) {
	testCases := []struct {
		Location       LocationModel
		ExpectedResult string
	}{
		{LocationModel{13.388860, 52.517037}, "13.388860,52.517037"},
		{LocationModel{13.38886, 52.51703700}, "13.388860,52.517037"},
		{LocationModel{13, 52}, "13.000000,52.000000"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("ShouldReturn`%s`From`%+v`", tc.ExpectedResult, tc.Location), func(t *testing.T) {
			if tc.Location.Description() != tc.ExpectedResult {
				t.Fatalf("LocationModel.Description() returned `%+v` , expected `%+v`", tc.Location.Description(), tc.ExpectedResult)
			}
		})
	}
}
