package main

import "fmt"

/*
built-in append function

func append(s []T, vs ...T) []T

somewhat like C++ vector.push_back
*/

func main() {
    var s []int
    printSlice(s)

    s = append(s, 0)
    printSlice(s)

    // slice len, cap grows as we append
    s = append(s, 1)
    printSlice(s) // len=2 cap=2 [0 1]

    // add multiple elements at once
    s = append(s, 2, 3, 4)
    printSlice(s) // len=5 cap=6 [0 1 2 3 4] why 6??
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
