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
}

// Flusher is an optional interface for Writers to implement
// if they need to be flushed after writing all the records.
type Flusher interface {
	Flush() error
}

// Copy provides that ability to copy tabulized data
// from one format to another.
//
func Copy(w Writer, r Reader) error {
	for {
		record, err := r.Read()
		if err == io.EOF {
			if f, ok := w.(Flusher); ok {
				return f.Flush()
			}
			return nil
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
