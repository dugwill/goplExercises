//ex5-19
//The instructions were to use Panic and recover to write a fucntion that does
//not have a return statement, but still returns a non-zero value.
// I had some trouble with this one and while looking around on the net
//I found https://www.goinggo.net/2013/06/understanding-defer-panic-and-recover.html
//Mr. Kennedy wrote this great tutorial about defer, panic and recover
//using his ideas I was able to finish the excerise, and I understant this
//Topic much better.

package main

import (
	"fmt"
)

func main() {

	if err := noReturn(); err != nil {
		fmt.Printf(" Got a return err of: %v\n", err)
	}

}

// noReturn does not contain a return statement but does manage to
// communicate the err state after a panic to main.
// the defer function calls recover and gets the panic message
// It sets err to the Panic message and returns that to main

func noReturn() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	panic(fmt.Sprintf("There was a panic:  %v\n", err))

}
