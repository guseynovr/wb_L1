package main

import (
	"log"
	"os"
)

func main() {
	const usage = "Usage: \"string to be reversed\""

	if len(os.Args) != 2 {
		log.Fatal(usage)
	}
	println(reversed(os.Args[1]))
}

func reversed(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}
