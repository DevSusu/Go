module github.com/DevSusu/Go/golang.org/hello

go 1.16

require (
	github.com/DevSusu/Go/golang.org/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

replace github.com/DevSusu/Go/golang.org/greetings => ../greetings
