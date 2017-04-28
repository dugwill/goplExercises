//Excersize 7.2
// I really struggled with this excersize.  I could not understanf the
// Concept of wrapping a the original 'Writer' interface.
// I finally found an example that got me thorugh this
// https://play.golang.org/p/ssz2AKIj_y

package countingWriter

import (
	"fmt"
	"io"
)

// Create a wrapper for the new io.Writer
type ioWrapper struct {
	io.Writer
}

// bytecounter will always hold the number of bytes written using
// the newWriter
var byteCounter int64

//func (w ioWrapper) Write, overrides the io.writer Write func
func (w ioWrapper) Write(p []byte) (n int, err error) {
	fmt.Printf("%s\n", p)        //Prints the data to Stdout. Just for fun, not necessary
	byteCounter += int64(len(p)) //adds the written bytes to counter
	return len(p), err
}

// Counting writer takes an io.Writer and returns an new writer
// based on type ioWrapper and the address of the byte counter
// The use of type ioWrapper forces  the use of 'func (w ioWrapper) Write'
// Whenever newWriter is used as an io.Writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var newWriter = ioWrapper{w}
	return newWriter, &byteCounter
}
