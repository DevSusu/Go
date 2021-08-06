package main

import "fmt"

type Vertex struct {
    // same type can be omitted
    X, Y int
}

var (
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 1} // Y is 0
    v3 = Vertex{}
    pp = &Vertex{2, 3}
)

func main() {
    // prints {1 2}
    v := Vertex{1, 2}
    fmt.Println(v)
    
    v.X = 4
    fmt.Println(v)

    p := &v
    p.X = 1e9 // don't need something like p->x , (*p).X
    fmt.Println(v)

    // {1 2} {1 0} {0 0} &{2 3}
    fmt.Println(v1, v2, v3, pp)
}
