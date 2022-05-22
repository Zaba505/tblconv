package plugin

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSQLDriver_Txns(t *testing.T) {
	t.Run("should successfully be able to execute a query without returning any rows", func(subT *testing.T) {
		args := getHelperPluginCLI("execute", "--LastInsertId=1", "--RowsAffected=1")
		d := NewDriver(args[0], WithArgs(args[1:]...), WithEnv("GO_WANT_HELPER_PROCESS=1"))
		db := sql.OpenDB(d)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, tx.Commit)

		result, err := tx.ExecContext(ctx, "INSERT INTO t (hello) VALUES world")
		if !assert.Nil(subT, err) {
			return
		}

		lastInsertId, err := result.LastInsertId()
		if !assert.Nil(subT, err) || !assert.Equal(subT, int64(1), lastInsertId) {
			return
		}

		rowsAffected, err := result.RowsAffected()
		if !assert.Nil(subT, err) || !assert.Equal(subT, int64(1), rowsAffected) {
			return
		}
	})

	t.Run("should successfully be able to execute a query with return rows", func(subT *testing.T) {
		args := getHelperPluginCLI("query", "--Columns=HELLO", "--ColumnTypes=VARCHAR", "--TotalRows=10")
		d := NewDriver(args[0], WithArgs(args[1:]...), WithEnv("GO_WANT_HELPER_PROCESS=1"))
		db := sql.OpenDB(d)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, tx.Commit)

		rows, err := tx.QueryContext(ctx, "SELECT * FROM t")
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, rows.Close)

		cols, err := rows.Columns()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, len(cols)) {
			return
		}

		colTypes, err := rows.ColumnTypes()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, len(colTypes)) {
			return
		}

		vals := make([]string, 0, 10)
		for rows.Next() {
			var val sql.NullString
			if err := rows.Scan(&val); !assert.Nil(subT, err) {
				return
			}
			vals = append(vals, val.String)
		}
		if err := rows.Err(); !assert.Nil(subT, err) {
			return
		}

		if !assert.Equal(subT, 10, len(vals)) {
			return
		}
	})

	t.Run("should successfully be able to execute a prepared statement", func(subT *testing.T) {
		args := getHelperPluginCLI("execute", "--LastInsertId=1", "--RowsAffected=1")
		d := NewDriver(args[0], WithArgs(args[1:]...), WithEnv("GO_WANT_HELPER_PROCESS=1"))
		db := sql.OpenDB(d)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, tx.Commit)

		stmt, err := tx.PrepareContext(ctx, "INSERT INTO t (hello) VALUES ?")
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, stmt.Close)

		result, err := stmt.ExecContext(ctx, "world")
		if !assert.Nil(subT, err) {
			return
		}

		lastInsertId, err := result.LastInsertId()
		if !assert.Nil(subT, err) || !assert.Equal(subT, int64(1), lastInsertId) {
			return
		}

		rowsAffected, err := result.RowsAffected()
		if !assert.Nil(subT, err) || !assert.Equal(subT, int64(1), rowsAffected) {
			return
		}
	})

	t.Run("should successfully be able to query with a prepared statement", func(subT *testing.T) {
		args := getHelperPluginCLI("query", "--Columns=HELLO", "--ColumnTypes=VARCHAR", "--TotalRows=10")
		d := NewDriver(args[0], WithArgs(args[1:]...), WithEnv("GO_WANT_HELPER_PROCESS=1"))
		db := sql.OpenDB(d)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, tx.Commit)

		stmt, err := tx.PrepareContext(ctx, "SELECT * FROM t WHERE HELLO = ?")
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, stmt.Close)

		rows, err := stmt.QueryContext(ctx, "world")
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, rows.Close)

		cols, err := rows.Columns()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, len(cols)) {
			return
		}

		colTypes, err := rows.ColumnTypes()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, len(colTypes)) {
			return
		}

		vals := make([]string, 0, 10)
		for rows.Next() {
			var val sql.NullString
			if err := rows.Scan(&val); !assert.Nil(subT, err) {
				return
			}
			vals = append(vals, val.String)
		}
		if err := rows.Err(); !assert.Nil(subT, err) {
			return
		}

		if !assert.Equal(subT, 10, len(vals)) {
			return
		}
	})
}
