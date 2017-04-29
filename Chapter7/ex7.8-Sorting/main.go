// Original Work
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// gopl.io Exercise 7.8
// Modifications Copyright © 2017 Douglas Will
// License: https://creativecommons.org/licenses/by-sa/4.0/

// ******  This exercise is not complete  *******//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

//!+main
type Track struct {
	Title   string
	Artist  string
	Album   string
	Year    int
	Length  time.Duration
	drummer string
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s"), "bob"},
	{"Go", "Moby", "Moby", 1992, length("3m37s"), "Joe"},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s"), "Frank"},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s"), "Susie"},
}

var headings = []string{
	"Title",
	"Artist",
	"Album",
	"Year",
	"Length",
	"Drummer",
}

//length converts a string with time indicators to a duration
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// Enter input loop, q is the esc char
	for scanner.Text() != "q" {
		list(headings)
		fmt.Print("Select a Heading:")

		//Get input from user
		scanner.Scan()

		// check for esc char q
		if scanner.Text() != "q" {
			// convert the input string to a int
			if i, err := strconv.Atoi(scanner.Text()); err == nil {
				//normalize the input to 0 base counting
				i -= 1

				//Check for valid choice and call the sort func
				if i < len(headings) && i > 0 {
					headings = sortHeadings(headings, i)
				} else {
					fmt.Println("Invalid Choice\n")
				}
			}
		}
	}
}

//-main

// sortHeading takes a slice of strings and one element from that slice
// It re-orders the slice, placing the selection in the first position
func sortHeadings(headings []string, i int) []string {

	fmt.Println("Selection:", headings[i])

	// Check to see if the selection the first in the list
	if headings[i] == headings[0] {
		fmt.Println("Order is correct")
		return headings
	}

	// tmp holds the selection as the items are reordered
	tmp := headings[i]

	// Reorder slice. Iterate through the slic, moving items
	// before the selection to the right
	for index := i; index > 0; index-- {
		headings[index] = headings[index-1]
	}

	//Place the selection at the front of the slice
	headings[0] = tmp
	// return the reordered slice
	return headings
}

// list prints the headings list to Stdout along with the escape char q
func list(headings []string) {
	fmt.Printf("Table Headings\n")
	for index := 0; index < len(headings); index++ {
		fmt.Printf("%d. %s\n", index+1, headings[index])
	}
	fmt.Println("q to quit\n")
}
