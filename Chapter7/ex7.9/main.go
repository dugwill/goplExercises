// Original Work
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// gopl.io Exercise 7.9
// Modifications Copyright © 2017 Douglas Will
// License: https://creativecommons.org/licenses/by-sa/4.0/

// ***********  This Exersize is not finished   ********/

package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
	"time"
)

// added for exercise 7.9
//!+template

var tracklist = template.Must(template.New("tracklist").Parse(`

<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Lenght}}</td>
</tr>

`))

var tracklist1 = template.Must(template.New("tracklist1").Parse(`

<tr>
  <td>{{.0.Title}}</td>
  <td>{{.0.Artist}}</td>
  <td>{{.0.Album}}</td>
  <td>{{.0.Year}}</td>
  <td>{{.0.Lenght}}</td>
</tr>

`))

//!-template

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
	{"Ready 2 Go", "Martin Solveig", "Smash", 2020, length("4m24s")},
	{"Don't go", "Alicia Keys", "Bobby", 2006, length("8m36s")},
	{"Not Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var headings = []string{
	"Title",
	"Artist",
	"Album",
	"Year",
	"Length",
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
		printTracks(tracks)
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
				if i < len(headings) && i >= 0 {

					// Sort headings
					headings = sortHeadings(headings, i)

					// sort table based on headings order
					// how to suppy the sort order??
					// Modified the customSort struct to include the sorted headings
					// Modified the Anonymous Func in customSort to iterate through
					// headings in their sort order
					sort.Sort(customSort{tracks, headings, func(x, y *Track) bool {
						// counter to iterate through the heading in sorted order
						for i := 0; i < 5; i++ {

							// for each sort layer, this Switch/Case selectis and
							// returns the correct bool answer for the layer
							switch headings[i] {
							case "Title":
								{
									if x.Title != y.Title {
										return x.Title < y.Title
									}
								}
							case "Artist":
								{
									if x.Artist != y.Artist {
										return x.Artist < y.Artist
									}
								}
							case "Album":
								{
									if x.Album != y.Album {
										return x.Album < y.Album
									}
								}
							case "Year":
								{
									if x.Year != y.Year {
										return x.Year < y.Year
									}
								}
							case "Length":
								{
									if x.Length != y.Length {
										return x.Length < y.Length
									} //if --- I needed to add comments to track the closing brackets
								} // case Length
							} // switch
						} // for
						return false
					}}) // anonymous func, customSort, sort
				} else {
					fmt.Println("Invalid Choice\n")
				}
			} //ifAtoi
		} //if scanner
	} //for scanner
} //main
//!-main

//!+sortHeading
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

//!-sortHeading

//!+ list
// list prints the headings list to Stdout along with the escape char q
func list(headings []string) {
	fmt.Printf("Sort order\n")
	for index := 0; index < len(headings); index++ {
		fmt.Printf("%d. %s\n", index+1, headings[index])
	}
	fmt.Println("q to quit\n")
}

//!- list

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
	fmt.Println()

	for _, t := range tracks {
		if err := tracklist.Execute(os.Stdout, t); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("%v\n", tracks[0].Title)

	//if err := tracklist1.Execute(os.Stdout, tracks); err != nil {
	//	log.Fatal(err)
	//}

}

//!-printTracks

//!+customcode
// Modified customSort struct to inlcude headings
type customSort struct {
	t        []*Track
	headings []string
	less     func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

//!-customcode
