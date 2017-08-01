// Copyright (c) 2017 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// callees shows possible targets of the selected function call
// using golang.org/x/tools/cmd/guru.
func callees(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-scope", scope(args), "-modified", "callees", s.pos()))
}

// callers shows possible callers of the selected function
// using golang.org/x/tools/cmd/guru.
func callers(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-scope", scope(args), "-modified", "callers", s.pos()))
}

// callstack shows the path from the callgraph root to the selected function
// using golang.org/x/tools/cmd/guru.
func callstack(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-scope", scope(args), "-modified", "callstack", s.pos()))
}

type GuruDef struct {
	Objpos, Descr string
}

// definition shows declaration of selected identifier
// using golang.org/x/tools/cmd/guru.
func definition(s selection, args []string) {
	var gd GuruDef
	js := runWithStdin(s.archive(), "guru", "-json", "-modified", "definition", s.pos())
	if err := json.Unmarshal([]byte(js), &gd); err != nil {
		log.Fatalf("failed to unmarshal guru json: %v\n", err)
	}
	if err := plumbText(gd.Objpos); err != nil {
		fmt.Println(gd.Objpos)
		log.Fatalf("failed to plumb: %v\n", err)
	}
}

// describe describes the selected syntax: definition, methods, etc.
// using golang.org/x/tools/cmd/guru.
func describe(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-modified", "describe", s.pos()))
}

// godoc shows documentation for items in Go source code
// using github.com/zmb3/gogetdoc.
func godoc(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "gogetdoc", "-modified", "-pos", s.pos()))
}

// extract extracts statements to a new function/method
// using github.com/godoctor/godoctor.
func extract(s selection, args []string) {
	if len(args) < 1 {
		log.Fatal("Usage: A ex <name>")
	}
	pos := fmt.Sprintf("%d,%d", s.start, s.end-s.start)
	stdin := bytes.NewReader(s.body)
	code := runWithStdin(stdin, "godoctor", "-scope", ".", "-complete", "-file", s.filename(), "-pos", pos, "extract", args[0])
	if i := strings.Index(code, "\n"); i != -1 {
		code = code[i+1:]
	}
	writeBody(s.win, code)
	reloadShowAddr(s.win, s.start)
}

// freevars shows free variables of the selection
// using golang.org/x/tools/cmd/guru.
func freevars(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-modified", "freevars", s.sel()))
}

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

// implements shows the 'implements' relation for the selected type or method
// using golang.org/x/tools/cmd/guru.
func implements(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-modified", "-scope", scope(args), "implements", s.pos()))
}

// peers shows send/receive corresponding to selected channel op
// using golang.org/x/tools/cmd/guru.
func peers(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-modified", "-scope", scope(args), "peers", s.sel()))
}

// pointsto shows variables the selected pointer may point to
// using golang.org/x/tools/cmd/guru.
func pointsto(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-modified", "-scope", scope(args), "pointsto", s.sel()))
}

// referrers shows all refs to the entity denoted by selected identifier
// using golang.org/x/tools/cmd/guru.
func referrers(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-modified", "referrers", s.pos()))
}

// rename renames the selected identifier
// using golang.org/x/tools/cmd/gorename.
func rename(s selection, args []string) {
	if len(args) < 1 {
		log.Fatal("Usage: A rn <name>")
	}
	run("gorename", "-offset", s.pos(), "-to", args[0])
	reloadShowAddr(s.win, s.start)
}

// share uploads the selected code to play.golang.org
// and prints the generated URL.
func share(s selection, args []string) {
	body := bytes.NewReader(s.body[s.start:s.end])
	req, err := http.NewRequest("POST", "https://play.golang.org/share", body)
	if err != nil {
		log.Fatalf("cannot send snippet: %v", err)
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("cannot send snippet: %v", err)
	}
	defer rsp.Body.Close()
	id, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatalf("cannot send snippet: %v", err)
	}
	fmt.Printf("https://play.golang.org/p/%s\n", id)
}

// what shows basic information about the selected syntax node
// using golang.org/x/tools/cmd/guru.
func what(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-modified", "what", s.pos()))
}

// whicherrs shows possible values of the selected error variable
// using golang.org/x/tools/cmd/guru.
func whicherrs(s selection, args []string) {
	fmt.Println(runWithStdin(s.archive(), "guru", "-modified", "-scope", scope(args), "whicherrs", s.pos()))
}

func scope(args []string) string {
	if len(args) == 0 {
		return "."
	}
	return args[0]
}
