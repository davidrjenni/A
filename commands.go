// Copyright (c) 2017 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/cmd/guru/serial"
)

type gomodifytagsOutput struct {
	Start int      `json:"start"`
	End   int      `json:"end"`
	Lines []string `json:"lines"`
	Errs  []string `json:"errors"`
}

// addTags adds tags to the selected struct fields
// using github.com/fatih/gomodifytags.
func addTags(s selection, args []string) {
	if len(args) < 1 {
		log.Fatal(`Usage: A addtags <tags> [options]
<tags>:	comma-separated tags to add, e.g. json,xml
[options]:	options to add, e.g. 'json=omitempty'`)
	}
	arguments := []string{
		"-file", s.filename(), "-modified", "-format", "json", "-line", s.lineSel(), "-add-tags", args[0],
	}
	if len(args) > 1 {
		arguments = append(arguments, "-add-options", args[1])
	}
	buf := runWithStdin(s.archive(), "gomodifytags", arguments...)
	var out gomodifytagsOutput
	if err := json.Unmarshal([]byte(buf), &out); err != nil {
		log.Fatal(err)
	}
	if err := s.win.Addr("%d,%d", out.Start, out.End); err != nil {
		log.Fatal(err)
	}
	if _, err := s.win.Write("data", []byte(strings.Join(out.Lines, "\n")+"\n")); err != nil {
		log.Fatal(err)
	}
	showAddr(s.win, s.start)
	if len(out.Errs) != 0 {
		fmt.Fprintln(os.Stderr, strings.Join(out.Errs, "\n"))
	}
}

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

// definition shows declaration of selected identifier
// using golang.org/x/tools/cmd/guru.
func definition(s selection, args []string) {
	var gd serial.Definition
	js := runWithStdin(s.archive(), "guru", "-json", "-modified", "definition", s.pos())
	if err := json.Unmarshal([]byte(js), &gd); err != nil {
		log.Fatalf("failed to unmarshal guru json: %v\n", err)
	}
	if err := plumbText(gd.ObjPos); err != nil {
		fmt.Println(gd.ObjPos)
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
	showAddr(s.win, s.start)
}

type output struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Code  string `json:"code"`
}

// fillstruct fills a struct literal with default values
// using github.com/davidrjenni/reftools/cmd/fillstruct.
func fillstruct(s selection, args []string) {
	buf := runWithStdin(s.archive(), "fillstruct", "-modified", "-file", s.filename(), "-offset", fmt.Sprintf("%d", s.start), "-line", fmt.Sprintf("%d", s.startLine))
	var res []output
	if err := json.Unmarshal([]byte(buf), &res); err != nil {
		log.Fatal(err)
	}
	for _, out := range res {
		if err := s.win.Addr("#%d,#%d", out.Start, out.End); err != nil {
			log.Fatal(err)
		}
		if _, err := s.win.Write("data", []byte(out.Code)); err != nil {
			log.Fatal(err)
		}
	}
	showAddr(s.win, s.start)
}

// fillswitch fills a (type) switch statement with case statements.
// using github.com/davidrjenni/reftools/cmd/fillswitch.
func fillswitch(s selection, args []string) {
	buf := runWithStdin(s.archive(), "fillswitch", "-modified", "-file", s.filename(), "-offset", fmt.Sprintf("%d", s.start), "-line", fmt.Sprintf("%d", s.startLine))
	var res []output
	if err := json.Unmarshal([]byte(buf), &res); err != nil {
		log.Fatal(err)
	}
	for _, out := range res {
		if err := s.win.Addr("#%d,#%d", out.Start, out.End); err != nil {
			log.Fatal(err)
		}
		if _, err := s.win.Write("data", []byte(out.Code)); err != nil {
			log.Fatal(err)
		}
	}
	showAddr(s.win, s.start)
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
	showAddr(s.win, s.start)
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

type posShortener string

func newPosShortener() posShortener {
	wd, err := os.Getwd()
	if err != nil {
		return posShortener("")
	}
	return posShortener(wd)
}

// Do shortens the pos (of the form "file:line:col") by converting the
// file part to a relative path if it is shorter.
func (ps posShortener) do(pos string) string {
	if ps == "" {
		return pos
	}
	i := bytes.LastIndexByte([]byte(pos), ':')
	if i < 0 {
		return pos
	}
	i = bytes.LastIndexByte([]byte(pos[:i]), ':')
	if i < 0 {
		return pos
	}
	fpath, addr := pos[:i], pos[i:]

	rel, err := filepath.Rel(string(ps), fpath)
	if err != nil || len(rel) > len(fpath) {
		return pos
	}
	return rel + addr
}

// referrers shows all refs to the entity denoted by selected identifier
// using golang.org/x/tools/cmd/guru.
func referrers(s selection, args []string) {
	ps := newPosShortener()

	var init serial.ReferrersInitial
	js := runWithStdin(s.archive(), "guru", "-json", "-modified", "referrers", s.pos())
	dec := json.NewDecoder(strings.NewReader(js))
	if err := dec.Decode(&init); err != nil {
		log.Fatalf("failed to unmarshal ReferrersInitial: %v\n", err)
	}
	fmt.Printf("%v: references to %v\n", ps.do(init.ObjPos), init.Desc)
	for {
		var pkg serial.ReferrersPackage
		if err := dec.Decode(&pkg); err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("failed to unmarshal ReferrersPackage: %v\n", err)
		}
		for _, r := range pkg.Refs {
			fmt.Printf("%v: %v\n", ps.do(r.Pos), r.Text)
		}
	}
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

// rmTags removes tags  the selected struct fields
// using github.com/fatih/gomodifytags.
func rmTags(s selection, args []string) {
	if len(args) < 1 {
		log.Fatal(`Usage: A rmtags <tags> [options]
<tags>:	comma-separated tags to remove, e.g. json,xml
[options]:	options to remove, e.g. 'json=omitempty'`)
	}
	arguments := []string{
		"-file", s.filename(), "-modified", "-format", "json", "-line", s.lineSel(), "-remove-tags", args[0],
	}
	if len(args) > 1 {
		arguments = append(arguments, "-remove-options", args[1])
	}
	buf := runWithStdin(s.archive(), "gomodifytags", arguments...)
	var out gomodifytagsOutput
	if err := json.Unmarshal([]byte(buf), &out); err != nil {
		log.Fatal(err)
	}
	if err := s.win.Addr("%d,%d", out.Start, out.End); err != nil {
		log.Fatal(err)
	}
	if _, err := s.win.Write("data", []byte(strings.Join(out.Lines, "\n")+"\n")); err != nil {
		log.Fatal(err)
	}
	showAddr(s.win, s.start)
	if len(out.Errs) != 0 {
		fmt.Fprintln(os.Stderr, strings.Join(out.Errs, "\n"))
	}
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
