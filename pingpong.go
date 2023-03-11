package main

import (
    "fmt"
    "time"
)

func ping(c1, c2 chan string) {
    for {
        msg := <-c1
        fmt.Println(msg)
        time.Sleep(1 * time.Second) // intervalo de 1 segundo
        c2 <- "pong"
    }
}

func pong(c1, c2 chan string) {
    for {
        msg := <-c1
        fmt.Println(msg)
        time.Sleep(1 * time.Second) // intervalo de 1 segundo
        c2 <- "ping"
    }
}

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go ping(c1, c2)
    go pong(c2, c1)

    c1 <- "ping"

    select {}
}
