/*
Copyright © 2021 Zaba505

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package excel

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
