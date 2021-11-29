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
	"errors"
	"io"
	"testing"
)

type badReader struct {
	err error
}

func (r badReader) Read() ([]string, error) {
	return nil, r.err
}

type badWriter struct {
	err error
}

func (w badWriter) Write(_ []string) error {
	return w.err
}

type badFlusher struct {
	Writer
	err error
}

func (w badFlusher) Flush() error {
	return w.err
}

func TestCopy(t *testing.T) {
	readErr := errors.New("readErr")
	writeErr := errors.New("writeErr")
	flushErr := errors.New("flushErr")

	testCases := []struct {
		Name        string
		Reader      Reader
		Writer      Writer
		ExpectedErr error
	}{
		{
			Name:   "WithFlusher",
			Reader: badReader{err: io.EOF},
			Writer: badFlusher{
				Writer: badWriter{},
				err:    flushErr,
			},
			ExpectedErr: flushErr,
		},
		{
			Name:        "WithoutFlusher",
			Reader:      badReader{err: io.EOF},
			Writer:      badWriter{},
			ExpectedErr: nil,
		},
		{
			Name:        "WithBadReader",
			Reader:      badReader{err: readErr},
			Writer:      badWriter{},
			ExpectedErr: readErr,
		},
		{
			Name:        "WithBadWriter",
			Reader:      badReader{},
			Writer:      badWriter{err: writeErr},
			ExpectedErr: writeErr,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(subT *testing.T) {
			err := Copy(testCase.Writer, testCase.Reader)
			if testCase.ExpectedErr != err {
				subT.Log(err)
				subT.Fail()
				return
			}
		})
	}
}
