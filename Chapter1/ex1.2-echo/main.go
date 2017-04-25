// Original Work
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// gopl.io Exercise 1.2
// Modifications Copyright © 2017 Douglas Will
// License: https://creativecommons.org/licenses/by-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	//DW - Modify the counter, i, to count from 0
	for i := 0; i < len(os.Args); i++ {
		//Modify to print each arg on a separate line
		fmt.Println(os.Args[i])
	}
}

//!-