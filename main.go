// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

const usage = `Usage: A <cmd>

Commands:
	def	shows declaration of selected identifier
	doc	shows documentation for items in Go source code
	ex	extracts statements to a new function/method
	impl	generate method stubs for implementing an interface
	rn	renames the selected identifier
	share	uploads the selected code to play.golang.org
`

var cmds = map[string]func(selection, []string){
	"def":   definition,
	"doc":   godoc,
	"ex":    extract,
	"impl":  impl,
	"rn":    rename,
	"share": share,
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatal(usage)
	}

	s, err := readSelection()
	if err != nil {
		log.Fatalf("cannot read selection: %v\n", err)
	}

	f, ok := cmds[os.Args[1]]
	if !ok {
		log.Fatal(usage)
	}
	f(s, os.Args[2:])
}

func run(cmd string, args ...string) string {
	return runWithStdin(nil, cmd, args...)
}

func runWithStdin(stdin io.Reader, cmd string, args ...string) string {
	var buf bytes.Buffer
	c := exec.Command(cmd, args...)
	c.Stderr = os.Stderr
	c.Stdout = &buf
	c.Stdin = stdin
	if err := c.Run(); err != nil {
		log.Fatalf("%s failed: %v\n", cmd, err)
	}
	return buf.String()
}

func archive(s selection) io.Reader {
	archive := fmt.Sprintf("%s\n%d\n%s", s.filename, len(s.body), string(s.body))
	return strings.NewReader(archive)
}
