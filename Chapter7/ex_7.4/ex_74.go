package ex_74

import (
	//"io"
	"strings"
)

type StringReader int

func (sr StringReader) NewReader(s string) *strings.Reader {

	readFromString := strings.NewReader(s)

	return (readFromString)
}
