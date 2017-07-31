// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"
	"testing"
)

type offsetTest struct {
	data       string
	offset     int
	byteOffset int
	line       int
}

var offsetTests = []offsetTest{
	{"abcdef", 0, 0, 1},
	{"abcdef", 1, 1, 1},
	{"abcdef", 5, 5, 1},
	{"日本語def", 0, 0, 1},
	{"日本語def", 1, 3, 1},
	{"日本語def", 5, 11, 1},
	{"日本語def\n", 7, 13, 2},
	{"日本語def\n日本語def", 13, 25, 2},
	{"日本語def\n日本語def\nabc", 17, 29, 3},
}

func TestByteOffset(t *testing.T) {
	for i, test := range offsetTests {
		off, line, err := byteOffset(strings.NewReader(test.data), test.offset)
		if err != nil {
			t.Errorf("%d: got error %v", i, err)
		}
		if off != test.byteOffset {
			t.Errorf("%d: expected byte offset %d, got %d", i, test.byteOffset, off)
		}
		if line != test.line {
			t.Errorf("%d: expected line number %d, got %d", i, test.line, line)
		}
	}
}
