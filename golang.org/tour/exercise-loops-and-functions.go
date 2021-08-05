package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    /*
    let's implement a square root function:
        given a number x, we want to find the number z for which z² is most nearly x.
    */

    z := 1.0
    
    // method 1: Sqrt(2) ~= 1.42
    // 0.41 accuracy
    // for ; (x - z*z) > 0.01; z += 0.01 {
    // }

    // method 2: Sqrt(2) = 1.414213562373095
    // ~ 100% accuracy
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2*z)
    }

    return z
}

func main() {
    // x - z*z > 0.01; z += 0.01  -->  Sqrt(2) = 1.42
    es := Sqrt(2)
    ans := math.Sqrt(2)
    
    fmt.Println(ans, es)
    fmt.Printf("%.2f\n", math.Abs(((es - ans)/ans)) * 100)
}
