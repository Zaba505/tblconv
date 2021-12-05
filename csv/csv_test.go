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

package csv

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Zaba505/tblconv"
)

func TestReader(t *testing.T) {
	testData := `hello,goodbye
world,world
`

	r := NewReader(strings.NewReader(testData))
	w := tblconv.NewRecordsWriter()

	err := tblconv.Copy(w, r)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestWriter(t *testing.T) {
	var b bytes.Buffer

	r := tblconv.NewRecordsReader()
	w := NewWriter(&b)

	err := tblconv.Copy(w, r)
	if err != nil {
		t.Error(err)
		return
	}
}
