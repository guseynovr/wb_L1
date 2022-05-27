package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	const usage = "Usage: number(int64) position(0-63) value(0|1)"
	if len(os.Args) != 4 {
		log.Fatal(usage)
	}
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(usage)
	}
	i, err := strconv.ParseInt(os.Args[2], 10, 8)
	if err != nil || i < 0 || i > 63 {
		log.Fatal(usage)
	}
	val, err := strconv.ParseInt(os.Args[3], 2, 2)
	if err != nil {
		log.Fatal(err, usage)
	}
	println("Before:\t", strconv.FormatInt(n, 2))
	if val == 1 {
		n = n | (1 << i)
	} else {
		// println(strconv.FormatInt(n, 2), "^", strconv.FormatInt(1<<i, 2))
		n = n &^ (1 << i)
	}
	println("After:\t", strconv.FormatInt(n, 2))
}
