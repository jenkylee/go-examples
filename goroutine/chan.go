package main

import (
    "fmt"
    "time"
)

func hello(ch chan int) {
    time.Sleep(2 * time.Second)
    fmt.Println("after 2 second hello!!!")
 
    ch <- 1000
}

func main() {
    ch := make(chan int)

    go hello(ch)

    fmt.Println("start hello, wait ...")

    v := <-ch

    fmt.Println("receive:", v)
}
