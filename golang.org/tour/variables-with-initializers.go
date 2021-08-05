package main

import "fmt"

// initialization with type and value
var i, j int = 1, 2

func main() {
    // type is guessed from the initial value
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}
