package tblconv

import (
	"testing"
)

func TestGetCellId(t *testing.T) {
	testCases := []struct {
		Row      int
		Col      int
		Expected string
	}{
		{
			Row:      1,
			Col:      1,
			Expected: "A1",
		},
		{
			Row:      5,
			Col:      1,
			Expected: "A5",
		},
		{
			Row:      1,
			Col:      27,
			Expected: "AA1",
		},
		{
			Row:      1,
			Col:      53,
			Expected: "BA1",
		},
		{
			Row:      1,
			Col:      28,
			Expected: "AB1",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Expected, func(subT *testing.T) {
			actual := getCellId(testCase.Row, testCase.Col)
			if testCase.Expected != actual {
				subT.Log(len(actual), actual)
				subT.Fail()
				return
			}
		})
	}
}
