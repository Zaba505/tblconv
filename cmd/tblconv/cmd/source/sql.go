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

package source

import (
	"database/sql"
	"io"

	"github.com/Zaba505/tblconv"

	"github.com/spf13/cobra"
)

var (
	server string
	query  string
	args   []string
	dsn    string
)

func init() {
	register(
		"sql",
		"Read data from a SQL database.",
		func(cmd *cobra.Command) {
			supportedServers := sql.Drivers()
			s := ""
			for i, ss := range supportedServers {
				s += ss
				if i < len(supportedServers)-1 {
					s += ", "
				}
			}

			cmd.Flags().StringVarP(&server, "sql-server", "s", "", "SQL server (possible values: "+s+")")
			cmd.Flags().StringVarP(&query, "query", "q", "", "SQL query for retrieving data")
			cmd.Flags().StringSliceVarP(&args, "arg", "a", []string{}, "Values for filling in query placeholder parameters")
			cmd.Flags().StringVar(&dsn, "dsn", "", "Database endpoint")

			cmd.MarkFlagRequired("sql-server")
			cmd.MarkFlagRequired("query")
			cmd.MarkFlagRequired("dsn")
		},
		func(_ io.Reader, cmd *cobra.Command) tblconv.Reader {
			db, err := sql.Open(server, dsn)
			if err != nil {
				panic(err)
			}

			return tblconv.NewSQLReader(db, query, interfaceSlicize(args)...)
		},
	)
}

func interfaceSlicize(ss []string) []interface{} {
	is := make([]interface{}, len(ss))
	for i := range ss {
		is[i] = ss[i]
	}
	return is
}
