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

package tblconv

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCopy(t *testing.T) {
	testCases := []struct {
		Name     string
		Reader   Reader
		Writer   func(io.Writer) Writer
		Expected io.Reader
	}{
		{
			Name: "CSVtoCSV",
			Reader: NewCSVReader(strings.NewReader(`hello,goodbye
world,world
`)),
			Writer: func(w io.Writer) Writer { return NewCSVWriter(w) },
			Expected: strings.NewReader(`hello,goodbye
world,world
`),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(subT *testing.T) {
			ex, err := ioutil.ReadAll(testCase.Expected)
			if err != nil {
				subT.Error(err)
				return
			}

			var out bytes.Buffer
			err = Copy(testCase.Writer(&out), testCase.Reader)
			if err != nil {
				subT.Error(err)
				return
			}

			if !bytes.Equal(ex, out.Bytes()) {
				subT.Log(out.Bytes())
				subT.Fail()
				return
			}
		})
	}
}
