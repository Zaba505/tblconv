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

package output

import (
	"fmt"
	"io"

	"github.com/Zaba505/tblconv"

	"github.com/spf13/cobra"
)

type withFlags func(*cobra.Command)

type withWriter func(io.Writer, *cobra.Command) tblconv.Writer

type out struct {
	name      string
	short     string
	withFlags withFlags
	writer    withWriter
}

var outMap = map[string]out{}

func register(name, short string, f withFlags, g withWriter) {
	outMap[name] = out{
		name:      name,
		short:     short,
		withFlags: f,
		writer:    g,
	}
}

func Commands() []*cobra.Command {
	cmds := make([]*cobra.Command, 0, len(outMap))

	for _, o := range outMap {
		cmd := &cobra.Command{
			Use:   fmt.Sprintf("%s [flags] FILE|-", o.name),
			Short: o.short,
		}

		o.withFlags(cmd)

		cmd.Flags().StringP("output", "o", "", "Filename to write data to.")

		cmds = append(cmds, cmd)
	}

	return cmds
}

func Writer(name string, w io.Writer, cmd *cobra.Command) tblconv.Writer {
	o, ok := outMap[name]
	if !ok {
		panic("tblconv: unknown output format: " + name)
	}

	return o.writer(w, cmd)
}
