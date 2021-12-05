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
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/xuri/excelize/v2"
)

// DefaultSheetName
var DefaultSheetName = "Sheet1"

type config struct {
	sheet string
}

// Option
type Option func(*config)

// SheetName
func SheetName(sheet string) Option {
	return func(cfg *config) {
		cfg.sheet = sheet
	}
}

// Reader
type Reader struct {
	cfg  config
	open func() (*excelize.File, error)

	idx  int
	rows [][]string
}

// NewReader
func NewReader(r io.Reader, opts ...Option) *Reader {
	cfg := config{
		sheet: "Sheet1",
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return &Reader{
		cfg:  cfg,
		open: func() (*excelize.File, error) { return excelize.OpenReader(r) },
	}
}

// Read
func (r *Reader) Read() ([]string, error) {
	if r.rows == nil {
		f, err := r.open()
		if err != nil {
			return nil, err
		}

		r.rows, err = f.GetRows(r.cfg.sheet)
		if err != nil {
			return nil, err
		}
	}

	if r.idx == len(r.rows) {
		return nil, io.EOF
	}
	idx := r.idx
	r.idx += 1
	return r.rows[idx], nil
}

// Writer
type Writer struct {
	flushOnce sync.Once

	cfg   config
	out   io.Writer
	excel *excelize.File

	idx int
}

// NewWriter
func NewWriter(w io.Writer, opts ...Option) *Writer {
	cfg := config{
		sheet: "Sheet1",
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	f := excelize.NewFile()

	hasSheet := false
	sheets := f.GetSheetList()
	for _, sheet := range sheets {
		if sheet == cfg.sheet {
			hasSheet = true
			break
		}
	}

	if !hasSheet {
		idx := f.NewSheet(cfg.sheet)
		f.SetActiveSheet(idx)
	}

	return &Writer{
		cfg:   cfg,
		out:   w,
		excel: excelize.NewFile(),
	}
}

// Write
func (w *Writer) Write(record []string) error {
	for i, val := range record {
		cellId := getCellId(w.idx+1, i+1)
		w.excel.SetCellStr(w.cfg.sheet, cellId, val)
	}
	w.idx += 1
	return nil
}

func getCellId(row, col int) string {
	colName := ""
	for col > 0 {
		modulo := (col - 1) % 26
		colName = string(rune(modulo+65)) + colName
		col = (col - modulo) / 26
	}
	return fmt.Sprintf("%s%d", strings.TrimSpace(colName), row)
}

// Flush
func (w *Writer) Flush() (err error) {
	w.flushOnce.Do(func() {
		_, err = w.excel.WriteTo(w.out)
	})
	return
}
