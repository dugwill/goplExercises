package ex_75

import (
	"io"
	"strings"
)

type LimReader int

func (lm LimReader) LimitReader(r io.Reader, n int64) io.Reader {

	var p []byte

	i := int(n)

	io.ReadAtLeast(r, p, i)

	str := string(p)

	return (strings.NewReader(str))

}
