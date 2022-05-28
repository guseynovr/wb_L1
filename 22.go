package main

import (
	"log"
	"math/big"
	"os"
	"strings"
)

const (
	prec = 1024
	base = 10
)

func main() {
	const usage = "Usage: number operation(+|-|*|/) number"

	if len(os.Args) != 4 {
		log.Fatal(usage)
	}
	x, _, err := big.ParseFloat(os.Args[1], base, prec, big.ToNearestEven)
	if err != nil {
		log.Fatal(err, usage)
	}
	y, _, err := big.ParseFloat(os.Args[3], base, prec, big.ToNearestEven)
	if err != nil {
		log.Fatal(err, usage)
	}
	println(x.Text('f', 10), y.Text('f', 10))
	op := os.Args[2]
	if !strings.ContainsAny(op, "+-*/") {
		log.Fatal(usage)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	var z big.Float
	switch op {
	case "+":
		z.Add(x, y)
	case "-":
		z.Sub(x, y)
	case "*":
		z.Mul(x, y)
	case "/":
		z.Quo(x, y)
	}

	println("result =", z.Text('f', 10))
}
