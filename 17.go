package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	size = 20
	rng  = 50
	value = 15
)

func main() {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, size)
	for i := range nums {
		nums[i] = rand.Intn(rng)
	}
	sort.Ints(nums)
	fmt.Println(nums)
	index := binarySearch(nums, value)
	if index < 0 {
		fmt.Printf("%d not found\n", value)	
	} else {
		fmt.Printf("Index of %d: %d\n", value, index)	
	}
}

func binarySearch(nums []int, val int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] < val {
			low = mid + 1
		} else if nums[mid] > val {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
