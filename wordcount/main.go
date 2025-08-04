package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	for _, w := range strings.Fields(s) {
		ret[w]++
	}

	return ret
}

func main() {
	s := "hello world"
	ret := WordCount(s)
	for w, c := range ret {
		fmt.Println(w, c)
	}
}
