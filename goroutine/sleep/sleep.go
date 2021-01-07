package sleep

import (
    "fmt"
    "time"
)

func hello() {
    time.Sleep(2 * time.Second)
    fmt.Println("after 2 second hello!!!")
}

func main() {
    go hello()

    fmt.Println("start hello, wait...")

    for {
       time.Sleep(1 * time.Second)
    }
}
