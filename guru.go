// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// callees shows possible targets of the selected function call
// using golang.org/x/tools/cmd/guru.
func callees(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-scope", scope(args), "-modified", "callees", s.pos()))
}

// callers shows possible callers of the selected function
// using golang.org/x/tools/cmd/guru.
func callers(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-scope", scope(args), "-modified", "callers", s.pos()))
}

// callstack shows the path from the callgraph root to the selected function
// using golang.org/x/tools/cmd/guru.
func callstack(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-scope", scope(args), "-modified", "callstack", s.pos()))
}

// describe describes the selected syntax: definition, methods, etc.
// using golang.org/x/tools/cmd/guru.
func describe(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-modified", "describe", s.pos()))
}

type GuruDef struct {
	Objpos, Descr string
}

// definition shows declaration of selected identifier
// using golang.org/x/tools/cmd/guru.
func definition(s selection, args []string) {
	var gd GuruDef
	js := runWithStdin(archive(s), "guru", "-json", "-modified", "definition", s.pos())
	if err := json.Unmarshal([]byte(js), &gd); err != nil {
		log.Fatalf("failed to unmarshal guru json: %v\n", err)
	}
	if err := plumbText(gd.Objpos); err != nil {
		fmt.Println(gd.Objpos)
		log.Fatalf("failed to plumb: %v\n", err)
	}
}

// freevars shows free variables of the selection
// using golang.org/x/tools/cmd/guru.
func freevars(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-modified", "freevars", s.sel()))
}

// implements shows the 'implements' relation for the selected type or method
// using golang.org/x/tools/cmd/guru.
func implements(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-modified", "-scope", scope(args), "implements", s.pos()))
}

// peers shows send/receive corresponding to selected channel op
// using golang.org/x/tools/cmd/guru.
func peers(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-modified", "-scope", scope(args), "peers", s.sel()))
}

// pointsto shows variables the selected pointer may point to
// using golang.org/x/tools/cmd/guru.
func pointsto(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-modified", "-scope", scope(args), "pointsto", s.sel()))
}

// referrers shows all refs to the entity denoted by selected identifier
// using golang.org/x/tools/cmd/guru.
func referrers(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-modified", "referrers", s.pos()))
}

// what shows basic information about the selected syntax node
// using golang.org/x/tools/cmd/guru.
func what(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-modified", "what", s.pos()))
}

// whicherrs shows possible values of the selected error variable
// using golang.org/x/tools/cmd/guru.
func whicherrs(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-modified", "-scope", scope(args), "whicherrs", s.pos()))
}

func scope(args []string) string {
	if len(args) == 0 {
		return "."
	}
	return args[0]
}
