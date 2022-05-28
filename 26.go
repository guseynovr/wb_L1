package main

import (
	"log"
	"os"
	"unicode"
)

func main() {
	const usage = "Usage: \"string to test\""

	if len(os.Args) != 2 {
		log.Fatal(usage)
	}
	println(testUnique(os.Args[1]))
}

func testUnique(s string) bool {
	runes := make(map[rune]struct{})
	for _, r := range s {
		r = unicode.ToLower(r)
		if _, exists := runes[r]; exists {
			return false
		}
		runes[r] = struct{}{}
	}
	return true
}
