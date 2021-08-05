package main

import (
    "fmt"
    "math"
)

func main () {
    // ./experted-names.go:9:17: cannot refer to unexported name math.pi
    // fmt.Println(math.pi)
    fmt.Println(math.Pi)
}
