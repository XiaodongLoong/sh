// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package interp

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/mvdan/sh/syntax"
)

func TestFile(t *testing.T) {
	cases := []struct {
		prog, out string
	}{
		{"", ""},
		{"echo foo", "foo\n"},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			file, err := syntax.Parse(strings.NewReader(c.prog), "", 0)
			if err != nil {
				t.Fatalf("could not parse: %v", err)
			}
			var buf bytes.Buffer
			r := Runner{
				File:   file,
				Stdout: &buf,
			}
			if err := r.Run(); err != nil {
				t.Fatalf("could not run: %v", err)
			}
			if got := buf.String(); got != c.out {
				t.Fatalf("unexpected output:\nwant: %q\ngot:  %q",
					c.out, got)
			}
		})
	}
}
