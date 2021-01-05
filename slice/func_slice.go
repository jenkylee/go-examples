package main

import "fmt"

func main() {
    s1 := []int{11, 22, 33, 44}
    foo(s1)
    fmt.Println(s1[1])
}

func foo(s []int) {
    for index, _ := range s {
        s[index] += 1
    }
}
