package main

import "fmt"

func main() {
    fmt.Println("counting")

    // defer calls are pushed to stack
    // so this counts DOWN
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
    
    // more explanation on using defer
    // https://blog.golang.org/defer-panic-and-recover
}
