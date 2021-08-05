package main

import "fmt"

/*
zero values are:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
*/

func main() {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
