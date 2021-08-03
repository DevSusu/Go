package main

import (
    "fmt"
    "log"

    "github.com/DevSusu/Go/golang.org/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0) // disable printing time, file, line number

    names := []string{"Subin", "Samantha", "Andy"}

    msgs, err := greetings.Hellos(names)
    if err != nil {
	log.Fatal(err)
    }

    fmt.Println(msgs)
}
