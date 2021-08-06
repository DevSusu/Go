package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    // loop through items in pow with index number
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

    // skip index, or value by using _
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
    for index, _ := range pow {
        fmt.Printf("%d\n", index)
    }
    for index := range pow { // equivalent
        fmt.Printf("%d\n", index)
    }
}
