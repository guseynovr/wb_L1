package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	size = 10
	rng  = 100
)

func main() {
	nums1 := make([]int, size)
	nums2 := make([]int, size)

	rand.Seed(time.Now().UnixNano())
	for i := range nums1 {
		nums1[i] = rand.Intn(rng)
		nums2[i] = rand.Intn(rng)
	}
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Print(intersection2(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]int)
	for _, x := range nums1 {
		set[x] = 1
	}
	for _, x := range nums2 {
		if _, ok := set[x]; ok {
			set[x] = 2
		}
	}
	result := []int{}
	j := 0
	for key := range set {
		if set[key] == 2 {
			result = append(result, key)
		}
		j++
	}
	return result
}

func intersection2(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	result := []int{}
	for _, x := range nums1 {
		set[x] = struct{}{}
	}
	for _, x := range nums2 {
		if _, ok := set[x]; ok {
			result = append(result, x)
		}
	}
	return result
}
