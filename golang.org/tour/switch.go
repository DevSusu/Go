package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")

    // has pre-statement just like if
    // actual switch is done on "os" variable
    switch os := runtime.GOOS; os {
	case "darwin": // + case doesn't need to be a constant
		fmt.Println("OS X.")
        // break isn't needed. switch only runs 1 case
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
