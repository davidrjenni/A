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
}

var offsetTests = []offsetTest{
	{"abcdef", 0, 0},
	{"abcdef", 1, 1},
	{"abcdef", 5, 5},
	{"日本語def", 0, 0},
	{"日本語def", 1, 3},
	{"日本語def", 5, 11},
}

func TestByteOffset(t *testing.T) {
	for _, test := range offsetTests {
		off, err := byteOffset(strings.NewReader(test.data), test.offset)
		if err != nil {
			t.Errorf("got error %v", err)
		}
		if off != test.byteOffset {
			t.Errorf("expected byte offset %d, got %d", test.byteOffset, off)
		}
	}
}
