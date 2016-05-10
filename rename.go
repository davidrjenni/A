// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "log"

// rename renames the selected identifier
// using golang.org/x/tools/cmd/gorename.
func rename(s selection, args []string) {
	if len(args) < 1 {
		log.Fatal("Usage: A rn <name>")
	}
	run("gorename", "-offset", s.pos(), "-to", args[0])
	reloadShowAddr(s.win, s.start)
}
