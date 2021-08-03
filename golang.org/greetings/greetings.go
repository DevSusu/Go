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
        return "", errors.New("empty name")
    }

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

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

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    formats := []string{
	"Hi, %v. Welcome!",
	"Great to see you, %v!",
	"Hail, %v! Well met!",
    }

    return formats[rand.Intn(len(formats))]
}
