// gopl.io Exercise 7.10
// License: https://creativecommons.org/licenses/by-sa/4.0/

// ***** This excercise is not finished 88888

package main

import (
	"fmt"
	"sort"
)

func main() {

	input := palitest{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}

	sort.Sort(sort.Reverse(input))
	fmt.Println(input)
	/*{
		fmt.Println(input + " is a Palindrome")
	}*/

}

func IsPalindrome(s sort.Interface) bool {

	return true
}

type palitest []int

func (x palitest) Len() int            { return len(x) }
func (x palitest) Less(i, j, int) bool { return i < j }
func (x palitest) Swap(i, j) int       { x[i], x[j] = x[j], x[i] }
