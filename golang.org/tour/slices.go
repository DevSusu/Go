package main

import "fmt"

func main() {
    // dynamically-sized, flexible view into the elements of an array
    // []T is a slice with elements with type T
    // a[low : high] half-open range [low, range)

    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[1:4]
    fmt.Println(s)

    s[0] = 100
    fmt.Println(primes)
    // values of primes is changed
    // s is just another pointer to primes's content

    // A slice does not store any data, it just describes a section of an underlying array.
    names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

    // slice literals
    // []int{1, 2, 3} creates an array, and then builds a slice
    q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
    fmt.Println(ss)

    s = primes[:]

    s = s[1:4]
    fmt.Println(s)

    s = s[:2] // s[0:2]
    fmt.Println(s)

    s = s[1:] // s[1:len]
    fmt.Println(s)
}
