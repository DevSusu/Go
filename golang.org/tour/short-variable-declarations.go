package main

import "fmt"

func main() {
    var i, j int = 1, 2

    // omitting "var" keyword
    // only possible inside functions
    // outside functions, "var" keyword is essential
    k := 3

    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}
