// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

const usage = `Usage: A <cmd>

Commands:
	share	uploads the selected code to play.golang.org
`

var cmds = map[string]func(selection, []string){
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
	var buf bytes.Buffer
	c := exec.Command(cmd, args...)
	c.Stderr = os.Stderr
	c.Stdout = &buf
	if err := c.Run(); err != nil {
		log.Fatalf("%s failed: %v\n", cmd, err)
	}
	return buf.String()
}
