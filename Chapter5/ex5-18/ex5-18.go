// Copyright © 2017 Douglas Will
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 148 of the Go Programming Lanugage.

//Exercise	5.18:
// the instruction were to re-write the fetch function to use and defer func
// to close the file 'f'.  Keeping the feature that would report the copy err, if any.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	//"golang.org/x/net/html"
)

var depth int

func main() {

	var filename string
	var size int64
	var err error

	for _, url := range os.Args[1:] {

		filename, size, err = fetch(url)
		if err != nil {
			log.Print(err)
		}

		fmt.Printf("Filename: %v\n\tFile Size: %v", filename, size)
	}
}

//!+
// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	// Defer closing the local file, but when closing, ensure
	// that the copy did not generate an error.  If it did
	// then do not over write the value of err
	defer func() {
		fmt.Println("closing file")
		if closeErr := f.Close(); err == nil { //close the file before checking the val of err
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
