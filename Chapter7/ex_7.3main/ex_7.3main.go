// This is a bit of code to test Excersize 7.3
// It generates 50 'randomd' numbers in a []int and passes
// the slice to sort.

package main

import (
	"fmt"
	"github.com/dugwill/gopl.io/ch7/ex_7.3"
	"math/rand"
)

func main() {
	data := make([]int, 50)
	fmt.Printf("Original list: ")
	for i := range data {
		data[i] = rand.Int() % 50
		fmt.Printf("%v ", data[i])
	}
	fmt.Println("")
	treesort.Sort(data)

}
