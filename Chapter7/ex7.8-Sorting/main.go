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
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var headings = []string{
	"Title",
	"Artist",
	"Album",
	"Year",
	"Length",
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Table Headings\n")

	list(headings)

	fmt.Println()

	fmt.Print("Select a Heading:")
	scanner.Scan()

	if i, err := strconv.Atoi(scanner.Text()); err == nil {

		i -= 1
		if i < len(headings) {

			fmt.Printf("You selected: %v\n", headings[i])

			if contains(headings, headings[i]) {
				fmt.Println("You made a valid choice.")
				headings = sortHeadings(headings, headings[i])
			}
		}
		fmt.Println("Invalid choice")
	}

}

func contains(strSlice []string, search string) bool {
	for _, value := range strSlice {
		if value == search {
			return true
		}
	}
	return false
}

// sortHeading takes a slice of strings and one element from that slice
// It re-orders the slice, placing the selection in the first position
func sortHeadings(headings []string, selection string) []string {

	fmt.Println("Selection:", selection)

	if selection == headings[0] {
		fmt.Println("Order is correct")
		return headings
	}
	// Reorder slice

	return headings
}

func list(headings []string) {
	for index := 0; index < len(headings); index++ {
		fmt.Printf("%d. %s\n", index+1, headings[index])
	}
}
