package main

import "fmt"

// same as defining x, y on top of function & returning that value
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x

    // "naked return"
    // only use in short functions like this, because of readability
    return
}

func main() {
    fmt.Println(split(17))
}
