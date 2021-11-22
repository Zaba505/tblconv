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
package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Zaba505/tblconv"
	"github.com/Zaba505/tblconv/cmd/tblconv/cmd/output"
	"github.com/Zaba505/tblconv/cmd/tblconv/cmd/source"

	"github.com/spf13/cobra"
)

func init() {
	for _, srcCmd := range source.Commands() {
		// every srcCmd needs there own cuz each cmd
		// can only have one parent
		intoCmd := &cobra.Command{
			Use:   "into",
			Short: "Convert between data formats.",
		}

		srcCmd.AddCommand(intoCmd)

		for _, outCmd := range output.Commands() {
			outCmd.Run = runConvert

			intoCmd.AddCommand(outCmd)
		}
	}
}

func runConvert(cmd *cobra.Command, args []string) {
	outputName, err := cmd.Flags().GetString("output")
	if err != nil {
		panic(err)
	}

	src := os.Stdin
	if args[0] != "-" {
		src, err = open(args[0])
		if err != nil {
			panic(err)
		}
	}

	dst := os.Stdout
	if strings.TrimSpace(outputName) != "" {
		dst, err = os.Create(outputName)
		if err != nil {
			panic(err)
		}
	}

	dstFmt := cmd.Name()
	srcFmt := cmd.Parent().Parent().Name()

	w := output.Writer(dstFmt, dst, cmd.Flags())
	r := source.Reader(srcFmt, src, cmd.Flags())

	err = tblconv.Copy(w, r)
	if err != nil {
		panic(err)
	}
}

func open(name string) (*os.File, error) {
	path, err := filepath.Abs(name)
	if err != nil {
		panic(err)
	}

	return os.Open(path)
}
