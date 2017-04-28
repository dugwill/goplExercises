// ContingWriterMain is a test program for Ex_7.2
// It passes an io.Writer to func countingWriter
// then prints to the new io.Writer and then prints
// the byteCounter

package main

import (
	"fmt"
	"github.com/dugwill/gopl.io/ch7/ex_7.2-countingWriter"
	"io"
)

func main() {

	var g io.Writer

	var countg *int64

	g, countg = countingWriter.CountingWriter(g)

	fmt.Println("Print to g")
	fmt.Fprintf(g, "Bob's your uncle")

	fmt.Printf("The count is g: %v\n", *countg)

	fmt.Println("Print to g again")
	fmt.Fprintf(g, "Now is the time")
	fmt.Printf("The count is g: %v\n", *countg)

}
