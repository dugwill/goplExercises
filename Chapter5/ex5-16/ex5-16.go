// Copyright © 2017 Douglas Will.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 143 in The GoPL book.

//ex5-16 Joins individual strings into a single string separated by a space	

//The instructions were to create a variadic function that joins
//strings similar to the golang strings.Join function

//the function joinStings take a variable list of strings
//joins them with a space sepatator and returns the new string

package main

import "fmt"

//!+
func joinStrings(words ...string) string {
	var list string
	for _, word := range words {
		list=list+" "+word
	}
	return list
}

//!-

func main() {
	//!+main
	fmt.Println("Joining word")
	fmt.Println(joinStrings())           //  "0"
	fmt.Println(joinStrings("Now"))          //  "3"
	fmt.Println(joinStrings("Now","is","the","time")) //  "10"

	//!+slice
	strSlice := []string{"Ask","not","what"}
	fmt.Println(joinStrings(strSlice...)) 
	//!-slice

	//!-main
}
