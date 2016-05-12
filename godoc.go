// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// godoc shows documentation for items in Go source code
// using github.com/zmb3/gogetdoc.
func godoc(s selection, args []string) {
	fmt.Println(runWithStdin(archive(s), "gogetdoc", "-pos", s.pos()))
}
