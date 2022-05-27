package main

import "fmt"

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, s := range strs {
		set[s] = struct{}{}
	}
	// fmt.Println(set)
	strs2 := []string{}
	for key := range set {
		strs2 = append(strs2, key)
	}
	fmt.Println(strs2)
}
