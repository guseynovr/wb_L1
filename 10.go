package main

import "fmt"

func main() {
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// fmt.Print(temps)
	result := make(map[int][]float32)
	for _, t := range temps {
		category := int(t) / 10 * 10
		result[category] = append(result[category], t)
	}
	fmt.Print(result)
}
