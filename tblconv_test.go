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
