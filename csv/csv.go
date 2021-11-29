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

// Package csv provides a Reader and Writer for the CSV format.
package csv

import (
	"encoding/csv"
	"io"
)

// NewReader
func NewReader(r io.Reader) *csv.Reader {
	return csv.NewReader(r)
}

// Writer
type Writer struct {
	CSV *csv.Writer
}

// NewWriter
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		CSV: csv.NewWriter(w),
	}
}

// Write
func (w *Writer) Write(record []string) error {
	return w.CSV.Write(record)
}

// Flush
func (w *Writer) Flush() error {
	w.CSV.Flush()
	return w.CSV.Error()
}
