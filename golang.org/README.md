# Learning Go

## Install
> https://golang.org/doc/install

```
# download and install

> echo "export PATH=/usr/local/go:$PATH" >> ~/.zshrc

# restart shell
> go version
go version go1.16.6 darwin/amd64
```

## Code

```
> mkdir golang.org && cd golang.org

> mkdir hello && cd hello
> go mod init github.com/DevSusu/Go/golang.org/hello

> cat go.mod
module github.com/DevSusu/Go/golang.org

go 1.16

> vi hello.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

> go run .
Hello, World!
```

## Package
> https://pkg.go.dev/

```
# add import statement
> vi hello.go
...
import "rsc.io/quote" # same as uri of [https://pkg.go.dev/rsc.io/quote]
...
    fmt.Println(quote.Go())
...

> go mod tidy
go: finding module for package rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c

# checksums added to go.sum
> cat go.sum

# package added to go.mod
> cat go.mod
...
require rsc.io/quote v1.5.2
```

## Create a Module

```
> mkdir greetings && cd greetings

> go mod init github.com/DevSusu/golang.org/greetings
go: creating new go.mod: module github.com/DevSusu/golang.org/greetings

> vi greetings.go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    // shortcut for declaring and assigning
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

> cd ../hello
> vi hello.go
package main

import (
    "fmt"

    "github.com/DevSusu/Go/golang.org/greetings"
)

func main() {
    // Get a greeting message and print it.
    msg := greetings.Hello("Subin")
    fmt.Println(msg)
}

# since greetings module isn't published
> go mod edit -replace github.com/DevSusu/Go/golang.org/greetings=../greetings
> cat go.mod
...
replace github.com/DevSusu/Go/golang.org/greetings => ../greetings

> go mod tidy
go: found github.com/DevSusu/Go/golang.org/greetings in github.com/DevSusu/Go/golang.org/greetings v0.0.0-00010101000000-000000000000

> cat go.mod
module github.com/DevSusu/Go/golang.org/hello

go 1.16

require (
	github.com/DevSusu/Go/golang.org/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

replace github.com/DevSusu/Go/golang.org/greetings => ../greetings

> go run .
Hi, Subin. Welcome!
```

## Handling Errors

```
> cd ../greetings
# 1. fix return values of Hello function
# 2. import errors package
# 3. check for blank names, return error if blank.
# 4. if not, return nil for error
> vi greetings.go
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}

> cd ../hello && vi hello.go
package main

import (
    "fmt"
    "log"

    "github.com/DevSusu/Go/golang.org/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0) // disable printing time, file, line number

    msg, err := greetings.Hello("")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(msg)
}

> go run .
greetings: empty name
exit status 1

# Slices
```
# declare slices with []type
# access slices with var[index]
> vi ../greetings/greetings.go
package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

> go run .
Hi, Subin. Welcome!
Great to see you, Subin!

# Maps

```
# 1. add Hellos function
# 2. loop through names (without index)
# 3. call Hello and add to msgs map
> vi ../greetings.go
func Hellos(names []string) (map[string]string, error) {
    msgs := make(map[string]string)
    for _, name := range names {
        msg, err := Hello(name)
        if err != nil {
            return nil, err
        }
        msgs[name] = msg
    }
    return msgs, nil
}

> vi hello.go
    names := []string{"Subin", "Samantha", "Andy"}
    msgs, err := greetings.Hellos(names)
    ...
    fmt.Println(msgs)

> go run .
map[Andy:Great to see you, Andy! Samantha:Hi, Samantha. Welcome! Subin:Hail, Subin! Well met!]
```

## Test
```
> vi ../greetings/greetings_test.go
package greetings

import (
    "testing"
    "regexp"
)

func TestHelloName(t *testing.T) {
    name := "Subin"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello(name)
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Subin) = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

> cd ../greetings
> go test
PASS
ok  	github.com/DevSusu/golang.org/greetings	0.005s

> go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok  	github.com/DevSusu/golang.org/greetings	0.005s
```

## Build

```
> cd ../hello
> go build
> ./hello

> go list -f '{{.Target}}'
/Users/susu/go/bin/hello

> export PATH=$PATH:/Users/susu/go/bin
> go install
> hello
map[Andy:Hail, Andy! Well met! Samantha:Great to see you, Samantha! Subin:Hail, Subin! Well met!]
```
