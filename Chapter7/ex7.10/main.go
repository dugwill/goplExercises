// gopl.io Exercise 7.10
// License: https://creativecommons.org/licenses/by-sa/4.0/

package main

import (
	"fmt"
)

func main() {

<<<<<<< HEAD
	fmt.Println(isPalindrome([]byte("racecar")))                   //true
	fmt.Println(isPalindrome([]byte("sore was I ere I saw eros"))) //true
	fmt.Println(isPalindrome([]byte("This is not a Palindrome")))  //false
=======
	input := palitest{7, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}

	sort.Sort(input)
	fmt.Println(input)

	input = palitest{7, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}

	sort.Reverse(input)
	fmt.Println(input)
>>>>>>> a724179ef410b1a9a142c7c756559f3991151419

}

func isPalindrome(s palitest) bool {
	//Setup a loop, i is index for left bytes, j is index for right bytes

	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) { //if i!<j and j!<i, then they must be equal
			continue //if equal continue
		} else {
			return false //If not equal then return false
		}
	}
	return true
}

type palitest []byte

func (x palitest) Len() int           { return len(x) }
func (x palitest) Less(i, j int) bool { return (x[i] < x[j]) }
func (x palitest) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x palitest) Len() int           { fmt.Println("Len"); return len(x) }
func (x palitest) Less(i, j int) bool { fmt.Println("Less"); return x[i] < x[j] } //This tests if i<j and returns t/f
func (x palitest) Swap(i, j int)      { fmt.Println("Swap"); x[i], x[j] = x[j], x[i] }
