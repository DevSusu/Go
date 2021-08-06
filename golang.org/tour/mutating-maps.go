package main

import "fmt"

func main() {
    m := make(map[string]int)

    // insert or update
    m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

    // delete key
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) // 0 default value

    // test if key is present. v is value, ok is true/false
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
