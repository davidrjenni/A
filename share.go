// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
