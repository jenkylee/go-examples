package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rr.r.Read(b)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		c := b[i]
		if c >= 'a' && c <= 'z' {
			if c > 'm' {
				b[i] = c - 13
			} else {
				b[i] = c + 13
			}
		} else if c >= 'M' && c <= 'Z' {
			if c > 'M' {
				b[i] = c - 13
			} else {
				b[i] = c + 13
			}
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Hello, World!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
