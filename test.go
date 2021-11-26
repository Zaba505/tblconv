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
	"io"
)

// RecordsReader is implemented solely for providing a simple reader to be used in tests.
type RecordsReader struct {
	records [][]string

	curRecord int
}

// NewRecordsReader
func NewRecordsReader(records ...[]string) *RecordsReader {
	return &RecordsReader{records: records}
}

// Read
func (r *RecordsReader) Read() ([]string, error) {
	if r.curRecord == len(r.records) {
		return nil, io.EOF
	}

	record := r.records[r.curRecord]
	r.curRecord += 1
	return record, nil
}

// RecordsWriter is implemented solely for providing a simple writer to be used in tests.
type RecordsWriter struct {
	records [][]string
}

// NewRecordsWriter
func NewRecordsWriter() *RecordsWriter {
	return new(RecordsWriter)
}

// Records returns all of the records currently written to the RecordsWriter.
func (w *RecordsWriter) Records() [][]string {
	return w.records
}

// Write
func (w *RecordsWriter) Write(record []string) error {
	w.records = append(w.records, record)
	return nil
}

// Flush
func (w *RecordsWriter) Flush() error {
	return nil
}
