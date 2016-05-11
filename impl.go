// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"strings"
)

// impl generates method stubs for implementing an interface
// using github.com/josharian/impl.
func impl(s selection, args []string) {
	if len(args) < 2 {
		log.Fatal("Usage: A impl <recv> <iface>")
	}
	l := len(args) - 1
	code := run("impl", strings.Join(args[0:l], " "), args[l:][0])
	end := ""
	if s.start < len(s.body)-1 {
		end = string(s.body[s.start+1:])
	}
	writeBody(s.win, string(s.body[:s.start])+"\n"+code+end)
	reloadShowAddr(s.win, s.start)
}
