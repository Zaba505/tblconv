package tblconv

import (
	"encoding/csv"
	"io"
)

// NewCSVReader
func NewCSVReader(r io.Reader) *csv.Reader {
	return csv.NewReader(r)
}

// CSVWriter
type CSVWriter struct {
	CSV *csv.Writer
}

// NewCSVWriter
func NewCSVWriter(w io.Writer) *CSVWriter {
	return &CSVWriter{
		CSV: csv.NewWriter(w),
	}
}

// Write
func (w *CSVWriter) Write(record []string) error {
	return w.CSV.Write(record)
}

// Flush
func (w *CSVWriter) Flush() error {
	w.CSV.Flush()
	return w.CSV.Error()
}
