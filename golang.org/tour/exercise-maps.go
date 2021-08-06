package main

import (
    "strings"
    "golang.org/x/tour/wc"
)

// Implement WordCount. It should return a map of the counts of each “word” in the string s
func WordCount(s string) map[string]int {
    m := map[string]int{}

    // string.Fields splits the string s around each instance of
    // one or more consecutive white space characters, as defined by unicode.IsSpace
    for _, word := range strings.Fields(s) {
        m[word]++
    }

    return m
}

func main() {
    wc.Test(WordCount)
}
