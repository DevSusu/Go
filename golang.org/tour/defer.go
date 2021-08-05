package main

import "fmt"

func main() {
    // called when surrounding function returns
    defer fmt.Println("world")

    fmt.Println("Hello")
}
