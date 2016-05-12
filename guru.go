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

