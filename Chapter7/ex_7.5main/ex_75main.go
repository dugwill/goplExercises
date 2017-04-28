
package main

import (
	"io"
	"log"
	"os"
	"strings"
	"github.com/dugwill/gopl.io/ch7/ex_75"
)

var lm ex_75.LimReader

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := .LimitReader(r, 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

}
