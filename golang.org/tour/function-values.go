package main

import (
    "fmt"
    "math"
)

/*
Functions are values too. They can be passed around just like other values.
Function values may be used as function arguments and return values.
*/

// compute function takes function(float64, float64) float64 and executes
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))

    fmt.Println(compute(hypot)) // sqrt(3*3 + 4*4) = 5
    fmt.Println(compute(math.Pow)) // 3^4 = 81
}
