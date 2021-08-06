package main

import "fmt"

func main() {
    var s []int
    fmt.Println(s, len(s), cap(s))

    // nil slice has 0 length, 0 capacity, and no underlying array
    if s == nil {
        fmt.Println("nil!")
    }
}
