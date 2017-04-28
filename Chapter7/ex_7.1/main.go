// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//!+bytecounter

type ByteCounter int

type WordCounter int // Create a new type 'WordCounter' based on type int

type LineCounter int // Create a new type 'LineCounter' based on type int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

// 'func (*WordCounter) Write' implements the Writer interface with a type *Wordcounter as a receiver
// the Writer interface is defined in the io package and included in fmt package
// Instead of formatting the output, it counts the words in the output,
// including the format string and the variable p, and returns the count
// The declaration needs to model the original declaration of Write
// Write(p [byte]) (n int, err, error), otherwise it would just be a standard func
func (w *WordCounter) Write(p []byte) (int, error) {
	pStr := string(p[:]) //Convert  []byte to string for use in reader
	scanner := bufio.NewScanner(strings.NewReader(pStr))

	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*w = WordCounter(count) // Set the value *w to the count, converting count to type WordCounter
	return count, nil

}

// 'func (*LineCounter) Write' implements the Writer interface with a type *LineCounter as a receiver
// the Writer interface is defined in the io package and included in fmt package
// Instead of formatting the output, it counts the lines in the output,
// including the format string and the variable p, and returns the count
// The declaration needs to model the original declaration of Write
// Write(p [byte]) (n int, err, error), otherwise it would just be a standard func
func (l *LineCounter) Write(p []byte) (int, error) {
	pStr := string(p[:])
	scanner := bufio.NewScanner(strings.NewReader(pStr))

	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*l = LineCounter(count) // Set the value *l to the count, converting count to type LineCounter
	return count, nil

}

//!-bytecounter

func main() {
	//!+main
	fmt.Println("Byte Counter")
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	c.Write([]byte("goodbye"))
	fmt.Println(c)

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s  Goodbye, %s\n", name, "Frank")
	fmt.Println(c) // "12", = len("hello, Dolly")
	fmt.Fprintf(&c, "Goodbye, %s", "Bob")
	fmt.Println(c)
	//!-main

	fmt.Fprintf(os.Stdout, "hello, %s  Goodbye, %s\n", name, "Frank")

	fmt.Println("Word Counter")
	var w WordCounter

	w.Write([]byte("Now is the time\nfor all good men\nto come to the aid\nof their country\n"))
	fmt.Println(w)

	w = 0
	var words = "for all good men"

	// when fmt.Fprintf is call here, the &w (type wordcount)
	// causes Fprintf to use the 'func (*WordCounter) Write' interface implementaion

	fmt.Fprintf(&w, "Now is the time\n%s to come to the aid\nof their country\n\n", words)
	fmt.Println(w)

	fmt.Println("Line Counter")
	var l LineCounter

	l.Write([]byte("Now is the time\nfor all good men\nto come to the aid\nof their country\n"))
	fmt.Println(l)

	l = 0
	var lines = "for all good men\n"
	// when fmt.Fprintf is call here, the &l (type wordcount)
	// causes Fprintf to use the 'func (*LineCounter) Write' interface implemenation

	fmt.Fprintf(&l, "Now is the time\n%s to come to the aid\nof their country\n\n", lines)
	fmt.Println(l)

}
