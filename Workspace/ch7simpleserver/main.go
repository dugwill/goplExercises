package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var someData data
	someData = "This is some data"
	log.Fatal(http.ListenAndServe("localhost:8000", someData))

}

type data string

func (someData data) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "%v", "<!DOCTYPE html>")
	fmt.Fprintf(w, "%v", "<html>")
	fmt.Fprintf(w, "%v", "<body>")
	fmt.Fprintf(w, "%v\n\n", someData)

	fmt.Fprintf(w, "%v", "<td><a href=\"http://localhost:8000\">Refresh</a>,</td>")

	fmt.Fprintf(w, "%v", "</body>")
	fmt.Fprintf(w, "%v", "</html>")
}
