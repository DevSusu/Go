package main

import "fmt"

// Go functions may be closures
// a function value that references variables from outside its body
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    // pos, neg itself are functions. funtion variables.
    // they each have their sum varaiable on their own
    pos, neg := adder(), adder()

    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i) // 0 1 3 ... 45,
            neg(-2*i), // 0 -2 -6 ... -90
        )
    }
}
