package tblconv

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSQLReader(t *testing.T) {
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

	r := NewSQLReader(db, query)
	w := NewRecordsWriter()

	err = Copy(w, r)
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

func TestSQLWriter(t *testing.T) {
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

	r := NewRecordsReader(records...)
	w := NewSQLWriter(db, query)

	err = Copy(w, r)
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
