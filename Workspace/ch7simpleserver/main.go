package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

var tracklist = template.Must(template.New("tracklist").Parse(`

<!DOCTYPE html>
<html>
<body>

<h1>{{"Tracklist"}}</h1>
<table>
<tr style='text-align: left'>
  <th><a href="http://localhost:8000/sortByTitle">Title</a></th>
  <th><a href="http://localhost:8000/sortByArtist">Artist</a></th>
  <th><a href="http://localhost:8000/sortByAlbum">Album</a></th>
  <th><a href="http://localhost:8000/sortByYear">Year</a></th>
  <th><a href="http://localhost:8000/sortByLength">Length</a></th> 
</tr>

{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>

</body>
</html>

`))

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type trackptr []*Track

var tracks = trackptr{
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

	log.Fatal(http.ListenAndServe("localhost:8000", tracks))

}

func (tracks trackptr) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if err := tracklist.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}

	for {
		switch req.URL.Path {
		case "/sortByTitle":
			headings = sortHeadings(headings, 0)
		case "/sortByArtist":
			headings = sortHeadings(headings, 1)
		case "/sortByAlbum":
			headings = sortHeadings(headings, 2)
		case "/sortByYear":
			headings = sortHeadings(headings, 3)
		case "/sortByLength":
			headings = sortHeadings(headings, 4)

		}
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
	}

	if err := tracklist.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}

}

//!+sortHeading
// sortHeading takes a slice of strings and one element from that slice
// It re-orders the slice, placing the selection in the first position
func sortHeadings(headings []string, i int) []string {

	//fmt.Println("Selection:", headings[i])

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
