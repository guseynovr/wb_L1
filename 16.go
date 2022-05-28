package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	size = 20
	rng  = 100
)

func main() {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, size)
	for i := range nums {
		nums[i] = rand.Intn(rng)
	}
	fmt.Println(nums)
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(nums []int, low, high int) {
	if low < high {
		p := partition(nums, low, high)
		quickSort(nums, low, p)
		quickSort(nums, p+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[(low+high)/2]
	for {
		for nums[low] < pivot {
			low++
		}
		for nums[high] > pivot {
			high--
		}
		if low >= high {
			return high
		}
		nums[low], nums[high] = nums[high], nums[low]
		low++
		high--
	}
}
