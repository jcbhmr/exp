// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build generate

//go:generate go run $GOFILE

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Download flags.go from the Go repository.
	response, err := http.Get("https://github.com/golang/go/raw/master/src/internal/goexperiment/flags.go")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status code %d", response.StatusCode)
	}
	f, err := os.Create("flags.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = io.Copy(f, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
