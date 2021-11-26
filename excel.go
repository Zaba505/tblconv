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
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/xuri/excelize/v2"
)

// DefaultSheetName
var DefaultSheetName = "Sheet1"

type excelConfig struct {
	sheet string
}

// ExcelOption
type ExcelOption func(*excelConfig)

// SheetName
func SheetName(sheet string) ExcelOption {
	return func(cfg *excelConfig) {
		cfg.sheet = sheet
	}
}

// ExcelReader
type ExcelReader struct {
	cfg  excelConfig
	open func() (*excelize.File, error)

	idx  int
	rows [][]string
}

// Read
func (r *ExcelReader) Read() ([]string, error) {
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

// NewExcelReader
func NewExcelReader(r io.Reader, opts ...ExcelOption) *ExcelReader {
	cfg := excelConfig{
		sheet: "Sheet1",
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return &ExcelReader{
		cfg:  cfg,
		open: func() (*excelize.File, error) { return excelize.OpenReader(r) },
	}
}

// ExcelWriter
type ExcelWriter struct {
	flushOnce sync.Once

	cfg   excelConfig
	out   io.Writer
	excel *excelize.File

	idx int
}

// Write
func (w *ExcelWriter) Write(record []string) error {
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
func (w *ExcelWriter) Flush() (err error) {
	w.flushOnce.Do(func() {
		_, err = w.excel.WriteTo(w.out)
	})
	return
}

// NewExcelWriter
func NewExcelWriter(w io.Writer, opts ...ExcelOption) *ExcelWriter {
	cfg := excelConfig{
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

	return &ExcelWriter{
		cfg:   cfg,
		out:   w,
		excel: excelize.NewFile(),
	}
}
