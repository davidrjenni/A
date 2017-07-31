// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"strings"
)

// extract extracts statements to a new function/method
// using github.com/godoctor/godoctor.
func extract(s selection, args []string) {
	if len(args) < 1 {
		log.Fatal("Usage: A ex <name>")
	}
	pos := fmt.Sprintf("%d,%d", s.start, s.end-s.start)
	code := run("godoctor", "-scope", ".", "-complete", "-file", s.filename, "-pos", pos, "extract", args[0])
	if i := strings.Index(code, "\n"); i != -1 {
		code = code[i+1:]
	}
	writeBody(s.win, code)
	reloadShowAddr(s.win, s.start)
}
