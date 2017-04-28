package main

import (
	"bufio"
	"fmt"
	"os"
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
	for _, h := range headings {
		fmt.Printf("     %v\n", h)
	}
	fmt.Println()

	fmt.Print("Select a Heading:")
	scanner.Scan()
	fmt.Printf("You selected: %v\n", scanner.Text())

	if contains(headings, scanner.Text) {
		fmt.Printf("You mand a valid choice")
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
