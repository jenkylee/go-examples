package main

import "fmt"

type MyReader struct {
}

func (r *MyReader) Read(b []byte) (n int, err error) {

	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

func main() {
	r := MyReader{}
	buf := make([]byte, 10)
	n, err := r.Read(buf)
	fmt.Println(n, err, string(buf))

	//reader.Validate(&MyReader{})
}
