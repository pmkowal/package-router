package models

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		location string
		expectedResult LocationModel
	}{
		{"13.388860,52.517037", LocationModel{13.388860, 52.517037}},
		{"13.388860", LocationModel{}},
		{"malformedString", LocationModel{}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("ShouldReturn`%+v`From`%s`", tc.expectedResult, tc.location), func(t *testing.T) {
			locationModel := LocationModel{}
			locationModel.Parse(tc.location)
			if locationModel != tc.expectedResult {
				t.Fatalf("Parse() returned `%+v` , expected `%+v`", locationModel, tc.expectedResult)
			}
		})
	}
}