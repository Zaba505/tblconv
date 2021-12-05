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

// Package csv contains a sql.Driver implementation for CSV.
package csv

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

func init() {
	sql.Register("csv", &Driver{})
}

// Driver
type Driver struct{}

// Open
func (d *Driver) Open(dsn string) (driver.Conn, error) {
	c, err := d.OpenConnector(dsn)
	if err != nil {
		return nil, err
	}
	return c.Connect(context.Background())
}

func (d *Driver) OpenConnector(name string) (driver.Connector, error) {
	return &connector{driver: d, fileName: name}, nil
}

type connector struct {
	driver *Driver

	fileName string
}

func (c *connector) Connect(ctx context.Context) (driver.Conn, error) {
	return &conn{}, nil
}

func (c *connector) Driver() driver.Driver {
	return c.driver
}

type conn struct{}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

func (c *conn) Close() error {
	return nil
}

func (c *conn) Begin() (driver.Tx, error) {
	return c.BeginTx(context.Background(), driver.TxOptions{})
}

func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return &tx{}, nil
}

type stmt struct{}

func (s *stmt) Close() error {
	return nil
}

func (s *stmt) NumInput() int {
	return 0
}

func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	return &result{}, nil
}

func (s *stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return &result{}, nil
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	return &rows{}, nil
}

func (s *stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return &rows{}, nil
}

type tx struct{}

func (tx *tx) Commit() error {
	return nil
}

func (tx *tx) Rollback() error {
	return nil
}

type result struct{}

func (r *result) LastInsertId() (int64, error) {
	return 0, nil
}

func (r *result) RowsAffected() (int64, error) {
	return 0, nil
}

type rows struct{}

func (r *rows) Columns() []string {
	return nil
}

func (r *rows) Close() error {
	return nil
}

func (r *rows) Next(dest []driver.Value) error {
	return nil
}
