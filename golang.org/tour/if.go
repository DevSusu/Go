package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
    // calc v before condition
    // v is only available in the if scope
    // () is not needed, but {} is required!
    if v := math.Pow(x, n); v < lim {
        return v
    }

    // can't mention v here
    return lim
}

func pow2(x, n, lim float64) float64 {
    // else statement
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }

    return lim
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
    fmt.Println(
        pow2(3, 2, 10),
        pow2(3, 3, 20),
    )

}
