package main

// Similar syntax as C
// Unlike C, Go has no pointer arithmetic.

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i
    fmt.Println(*p) // 42
    *p = 21
    fmt.Println(i) // value of i changed to 21

    p = &j
    *p = *p / 37
    fmt.Println(j) // 73
}
