package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

/*
Int to char conversions
97 = 'a'
109 = 'm'
64 = 'A'
77 = 'M'
*/

func (r2 rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r2.r.Read(b)
	for i := range b {
		if b[i] > 63 && b[i] < 78 {
			b[i] += 13
		} else if b[i] > 96 && b[i] < 110 {
			b[i] += 13
		} else {
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
