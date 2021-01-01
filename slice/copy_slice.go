package main

import "fmt"

func main() {
    s1 := []int{11, 22, 33}
    s2 := make([]int, 5)
    s3 := make([]int, 2)

    num := copy(s2, s1)
    copy(s3, s1)

    fmt.Println(num)
    fmt.Println(s2)
    fmt.Println(s3)
}
