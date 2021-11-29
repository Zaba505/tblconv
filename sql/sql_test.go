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
	"database/sql/driver"
	"testing"

	"github.com/Zaba505/tblconv"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestReader(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	records := [][]string{
		{"0", "tony", "stark", "32"},
		{"1", "clark", "kent", "2"},
		{"2", "bruce", "wayne", "-23"},
	}
	rows := sqlmock.NewRows([]string{"id", "first", "last", "age"})
	for i, record := range records {
		rows.AddRow(i, record[1], record[2], record[3])
	}

	query := "test reader"

	mock.ExpectBegin()
	mock.ExpectQuery(query).WillReturnRows(rows).RowsWillBeClosed()
	mock.ExpectCommit()

	r := NewReader(db, query)
	w := tblconv.NewRecordsWriter()

	err = tblconv.Copy(w, r)
	if err != nil {
		t.Error(err)
		return
	}

	// ensure all expectations have been met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Logf("unmet expectation error: %s", err)
		t.Fail()
		return
	}

	actualRecords := w.Records()
	if len(records) != len(actualRecords) {
		t.Logf("expected: %v\ngot: %v", records, actualRecords)
		t.Fail()
		return
	}

	for i, record := range records {
		for j, expected := range record {
			if expected != actualRecords[i][j] {
				t.Logf("expected: %s\ngot: %s", expected, actualRecords[i][j])
				t.Fail()
				return
			}
		}
	}
}

func TestWriter(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	records := [][]string{
		{"0", "tony", "stark", "32"},
		{"1", "clark", "kent", "2"},
		{"2", "bruce", "wayne", "-23"},
	}

	query := "INSERT ? ? ? ?"

	mock.ExpectBegin()
	for i, record := range records {
		mock.ExpectExec("^INSERT (.+) (.+) (.+) (.+)").
			WithArgs(convert2DriverValues(record)...).
			WillReturnResult(sqlmock.NewResult(int64(i), 1))
	}
	mock.ExpectCommit()

	r := tblconv.NewRecordsReader(records...)
	w := NewWriter(db, query)

	err = tblconv.Copy(w, r)
	if err != nil {
		t.Error(err)
		return
	}

	// ensure all expectations have been met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Logf("unmet expectation error: %s", err)
		t.Fail()
		return
	}
}

func convert2DriverValues(record []string) []driver.Value {
	vals := make([]driver.Value, 0, len(record))
	for _, val := range record {
		vals = append(vals, val)
	}
	return vals
}
