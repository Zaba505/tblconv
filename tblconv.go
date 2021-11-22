// Package tblconv
package tblconv

import (
	"io"
)

// Reader
type Reader interface {
	Read() ([]string, error)
}

// Writer
type Writer interface {
	Write(record []string) error
	Flush() error
}

// Copy provides that ability to copy tabulized data
// from one format to another.
//
func Copy(w Writer, r Reader) error {
	for {
		record, err := r.Read()
		if err == io.EOF {
			return w.Flush()
		}
		if err != nil {
			return err
		}

		err = w.Write(record)
		if err != nil {
			return err
		}
	}
}
