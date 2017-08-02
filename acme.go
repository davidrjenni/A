// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"9fans.net/go/acme"
	"9fans.net/go/plan9"
	"9fans.net/go/plumb"
)

type selection struct {
	win                *acme.Win
	start, end         int
	startLine, endLine int
	fname              string
	body               []byte
}

func (s selection) archive() io.Reader {
	return strings.NewReader(fmt.Sprintf("%s\n%d\n%s", s.filename(), len(s.body), s.body))
}

func (s selection) filename() string {
	if s.fname == "" {
		return "-"
	}
	return s.fname
}

func (s selection) pos() string {
	return fmt.Sprintf("%s:#%d", s.filename(), s.start)
}

func (s selection) sel() string {
	return fmt.Sprintf("%s,#%d", s.pos(), s.end)
}

func (s selection) lineSel() string {
	return fmt.Sprintf("%d,%d", s.startLine, s.endLine)
}

type bodyReader struct{ *acme.Win }

func (r bodyReader) Read(b []byte) (int, error) {
	return r.Win.Read("body", b)
}

type dataWriter struct{ *acme.Win }

func (w dataWriter) Write(data []byte) (int, error) {
	return w.Win.Write("data", data)
}

func readSelection() (s selection, err error) {
	id, err := strconv.Atoi(os.Getenv("winid"))
	if err != nil {
		return s, err
	}
	s.win, err = acme.Open(id, nil)
	if err != nil {
		return s, err
	}
	s.fname, err = readFilename(s.win)
	if err != nil {
		return s, err
	}
	s.body, err = ioutil.ReadAll(&bodyReader{s.win})
	if err != nil {
		return s, err
	}
	q0, q1, err := readAddr(s.win)
	if err != nil {
		return s, err
	}
	s.start, s.startLine, err = byteOffset(bytes.NewReader(s.body), q0)
	if err != nil {
		return s, err
	}
	s.end, s.endLine, err = byteOffset(bytes.NewReader(s.body), q1)
	if err != nil {
		return s, err
	}
	return
}

func readFilename(win *acme.Win) (string, error) {
	b, err := win.ReadAll("tag")
	if err != nil {
		return "", err
	}
	tag := string(b)
	if i := strings.Index(tag, " "); i != -1 {
		return tag[0:i], nil
	}
	return "", fmt.Errorf("cannot get filename from tag")
}

func readAddr(win *acme.Win) (q0, q1 int, err error) {
	if _, _, err := win.ReadAddr(); err != nil {
		return 0, 0, err
	}
	if err := win.Ctl("addr=dot"); err != nil {
		return 0, 0, err
	}
	return win.ReadAddr()
}

func byteOffset(r io.RuneReader, q int) (off, line int, err error) {
	line = 1 // the first line is line 1
	for i := 0; i != q; i++ {
		r, s, err := r.ReadRune()
		if err != nil {
			return 0, 0, err
		}
		if r == '\n' {
			line++
		}
		off += s
	}
	return
}

func reloadShowAddr(win *acme.Win, off int) error {
	if err := win.Ctl("get"); err != nil {
		return err
	}
	return showAddr(win, off)
}

func showAddr(win *acme.Win, off int) error {
	if err := win.Addr("#%d", off); err != nil {
		return err
	}
	return win.Ctl("dot=addr\nshow")
}

func writeBody(win *acme.Win, body string) error {
	if err := win.Ctl("nomark"); err != nil {
		fmt.Fprintf(os.Stderr, "failed to set nomark: %s", err)
	}
	defer func() {
		if err := win.Ctl("mark"); err != nil {
			fmt.Fprintf(os.Stderr, "failed to set mark: %s", err)
		}
	}()
	if err := win.Addr("0,$"); err != nil {
		return err
	}
	_, err := io.Copy(dataWriter{win}, strings.NewReader(body))
	return err
}

func plumbText(data string) error {
	f, err := plumb.Open("send", plan9.OWRITE)
	if err != nil {
		return err
	}
	defer f.Close()

	m := &plumb.Message{
		Src:  "A",
		Dst:  "edit",
		Dir:  "/",
		Type: "text",
		Data: []byte(data),
	}
	return m.Send(f)
}
