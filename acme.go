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
)

type selection struct {
	win        *acme.Win
	start, end int
	filename   string
	body       []byte
}

func (s selection) pos() string {
	return fmt.Sprintf("%s:#%d", s.filename, s.start)
}

func (s selection) sel() string {
	return fmt.Sprintf("%s,#%d", s.pos(), s.end)
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
	s.filename, err = readFilename(s.win)
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
	s.start, err = byteOffset(bytes.NewReader(s.body), q0)
	if err != nil {
		return s, err
	}
	s.end, err = byteOffset(bytes.NewReader(s.body), q1)
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

func byteOffset(r io.RuneReader, q int) (off int, err error) {
	for i := 0; i != q; i++ {
		_, s, err := r.ReadRune()
		if err != nil {
			return 0, err
		}
		off += s
	}
	return
}

func reloadShowAddr(win *acme.Win, off int) error {
	if err := win.Ctl("get"); err != nil {
		return err
	}
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
