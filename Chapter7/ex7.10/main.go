// gopl.io Exercise 7.10
// License: https://creativecommons.org/licenses/by-sa/4.0/

// ***** This excercise is not finished 88888

package main

import (
	"fmt"
	"sort"
)

func main() {

	input := []int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}

	if sort.Sort(sort.Reverse(sort.IntSlice(input))) {
		fmt.Println(input + " is a Palindrome")
	}

}

func IsPalindrome(s sort.Interface) bool {

	return true
}
