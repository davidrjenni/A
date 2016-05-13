// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// describe describes the selected syntax: definition, methods, etc.
// using golang.org/x/tools/cmd/guru.
func describe(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-modified", "describe", s.pos()))
}

// definition shows declaration of selected identifier
// using golang.org/x/tools/cmd/guru.
func definition(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "guru", "-modified", "definition", s.pos()))
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

func scope(args []string) string {
	if len(args) == 0 {
		return "."
	}
	return args[0]
}

