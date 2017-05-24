// gopl.io Exercise 7.10
// License: https://creativecommons.org/licenses/by-sa/4.0/

// ***** This excercise is not finished 88888

package main

import (
	"fmt"
	"sort"
)

func main() {

	input := palitest{7, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}

	sort.Sort(input)
	fmt.Println(input)

	input = palitest{7, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}

	sort.Reverse(input)
	fmt.Println(input)

}

func IsPalindrome(s sort.Interface) bool {

	return true
}

type palitest []int

func (x palitest) Len() int           { fmt.Println("Len"); return len(x) }
func (x palitest) Less(i, j int) bool { fmt.Println("Less"); return x[i] < x[j] } //This tests if i<j and returns t/f
func (x palitest) Swap(i, j int)      { fmt.Println("Swap"); x[i], x[j] = x[j], x[i] }
