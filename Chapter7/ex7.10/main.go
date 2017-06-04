// gopl.io Exercise 7.10
// License: https://creativecommons.org/licenses/by-sa/4.0/

package main

import (
	"fmt"
)

func main() {

	fmt.Println(isPalindrome([]byte("racecar")))                   //true
	fmt.Println(isPalindrome([]byte("sore was I ere I saw eros"))) //true
	fmt.Println(isPalindrome([]byte("This is not a Palindrome")))  //false

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
