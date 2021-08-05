package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()

    // equivalent as switch true
    // better way of writing long if-then-else chains
    switch {
    case t.Hour() < 12: // true == (t.Hour() < 12)
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon!")
    default:
        fmt.Println("Good evening.")
    }
}
