package main

import "fmt"

func main() {
    sum := 0
    // no () around, but {} is always required!!
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    sum = 1
    // init statement, post statements are optional
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)

    // while is spelled for in Go
    sum = 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)

    // these are inifinite loops
    for {}
    for {
    }
}
