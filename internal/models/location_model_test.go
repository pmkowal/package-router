package models

import (
	"fmt"
	"testing"
)

func TestLocationModel_Parse(t *testing.T) {
	testCases := []struct {
		location string
		expectedResult LocationModel
	}{
		{"13.388860,52.517037", LocationModel{13.388860, 52.517037}},
		{"13.388860", LocationModel{}},
		{"m13.388860,52.517037", LocationModel{}},
		{"malformedString", LocationModel{}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("ShouldReturn`%+v`From`%s`", tc.expectedResult, tc.location), func(t *testing.T) {
			locationModel := LocationModel{}
			locationModel.Parse(tc.location)
			if locationModel != tc.expectedResult {
				t.Fatalf("LocationModel.Parse() returned `%+v` , expected `%+v`", locationModel, tc.expectedResult)
			}
		})
	}
}

func TestLocationModel_Description(t *testing.T) {
	testCases := []struct {
		location LocationModel
		expectedResult string
	}{
		{LocationModel{13.388860, 52.517037}, "13.388860,52.517037"},
		{LocationModel{13.38886, 52.51703700}, "13.388860,52.517037"},
		{LocationModel{13, 52}, "13.000000,52.000000"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("ShouldReturn`%s`From`%+v`", tc.expectedResult, tc.location), func(t *testing.T) {
			if tc.location.Description() != tc.expectedResult {
				t.Fatalf("LocationModel.Description() returned `%+v` , expected `%+v`", tc.location.Description(), tc.expectedResult)
			}
		})
	}
}
