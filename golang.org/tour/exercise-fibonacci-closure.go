package main

import "fmt"

// Implement a fibonacci function that returns a function (a closure)
// that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a, b, val := 0, 1, 0

    return func() int {
        val = a
        c := a + b
        a = b
        b = c

        return val
    }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

