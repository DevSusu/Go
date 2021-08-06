package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

    s = s[:0]
    printSlice(s)

    s = s[:4] // extending length
    printSlice(s)

    s = s[2:]
    printSlice(s)

    // whatever direction you slice, length decreases
    // capacity only changes when slicing from front

    // index cannot be negative, so extending to front is impossible
    // s = s[-2:]
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
