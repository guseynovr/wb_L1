package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	const usage = "Usage: \"string to be reversed\""

	if len(os.Args) != 2 {
		log.Fatal(usage)
	}
	println(reversedWords(os.Args[1]))
}

func reversedWords(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}
