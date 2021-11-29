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
	"io"

	"github.com/Zaba505/tblconv"
	"github.com/Zaba505/tblconv/excel"

	"github.com/spf13/cobra"
)

func init() {
	register(
		"excel",
		"Read data formatted as CSV.",
		func(cmd *cobra.Command) {
			cmd.Flags().StringP("sheet", "s", excel.DefaultSheetName, "Excel sheet name to read values from.")
		},
		func(r io.Reader, cmd *cobra.Command) tblconv.Reader {
			sheet, err := cmd.Flags().GetString("sheet")
			if err != nil {
				panic(err)
			}
			return excel.NewReader(r, excel.SheetName(sheet))
		},
	)
}
