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

package sql

import (
	"context"
	"database/sql"
	"io"
	"time"
)

// Reader
type Reader struct {
	db     *sql.DB
	tx     *sql.Tx
	tctx   context.Context
	cancel func()

	query string
	args  []interface{}

	rows        *sql.Rows
	columnNames []string
}

// NewReader
func NewReader(db *sql.DB, query string, args ...interface{}) *Reader {
	return &Reader{
		db:    db,
		query: query,
		args:  args,
	}
}

// Read
func (r *Reader) Read() (record []string, err error) {
	if r.tx == nil {
		r.tctx, r.cancel = context.WithCancel(context.Background())
		r.tx, err = r.db.BeginTx(r.tctx, nil)
		if err != nil {
			return
		}

		r.rows, err = query(r.tctx, r.tx, r.query, r.args...)
		if err != nil {
			r.rollback()
			return
		}
	}

	if !r.rows.Next() {
		err = r.rows.Err()
		if err != nil {
			return
		}

		r.commitAndCloseRows()
		return nil, io.EOF
	}

	if r.columnNames == nil {
		r.columnNames, err = r.rows.Columns()
		if err != nil {
			r.commitAndCloseRows()
			return
		}
	}

	return scan(r.rows, r.columnNames)
}

func (r *Reader) rollback() {
	r.tx.Rollback()
	r.cancel()
	r.tx = nil
	r.tctx = nil
}

func (r *Reader) commitAndCloseRows() {
	r.rows.Close()
	r.tx.Commit()
	r.cancel()
	r.tx = nil
	r.tctx = nil
	r.rows = nil
}

func query(tctx context.Context, tx *sql.Tx, query string, args ...interface{}) (*sql.Rows, error) {
	return tx.QueryContext(tctx, query, args...)
}

func scan(rows *sql.Rows, columnNames []string) ([]string, error) {
	record := make([]string, len(columnNames))
	refs := make([]interface{}, 0, len(record))
	for i := range record {
		refs = append(refs, &record[i])
	}

	err := rows.Scan(refs...)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// Writer
type Writer struct {
	db *sql.DB
	tx *sql.Tx

	query string
}

// NewWriter
func NewWriter(db *sql.DB, query string) *Writer {
	return &Writer{
		db:    db,
		query: query,
	}
}

// Write uses the given record to fill an placeholder parameters in the query
// given when the SQLWriter was created.
//
// Writes occur within a sql.Tx and *DO NOT* periodically auto flush. Periodic
// flushing is the responsibility of the caller by utilizing SQLWriter.Flush().
// After flushing, a new sql.Tx will be created so with periodic flushing
// there is no gaurantee that all writes will occur in the same transaction.
//
func (w *Writer) Write(record []string) (err error) {
	if w.tx == nil {
		w.tx, err = w.db.Begin()
		if err != nil {
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := interfaceSlicize(record)
	_, err = w.tx.ExecContext(ctx, w.query, args...)
	return
}

func interfaceSlicize(ss []string) []interface{} {
	is := make([]interface{}, len(ss))
	for i := range ss {
		is[i] = ss[i]
	}
	return is
}

// Flush commits the underlying sql.Tx. See SQLWriter.Write() for more
// details about the relationship between Write and Flush for SQLWriter.
//
func (w *Writer) Flush() error {
	if w.tx == nil {
		return sql.ErrTxDone
	}
	tx := w.tx
	w.tx = nil
	return tx.Commit()
}
