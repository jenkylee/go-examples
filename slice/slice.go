package main

import "fmt"

func main() {
    slice := make([]int, 3, 5)

    fmt.Println(slice)
    fmt.Println(len(slice))
    fmt.Println(cap(slice))
    
    println(slice)
}
