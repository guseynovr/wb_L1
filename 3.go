package main

import (
	// "fmt"
	"sync"
	"sync/atomic"
)

// type Sum struct {
// 	value int64
// 	sync.Mutex
// }

func main() {
	// nums := []int{2, 4, 6, 8, 10}
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = (i + 1) * 2
	}
	// SquaresSumMutex(nums)
	SquaresSumAtomic(nums)
}

// SquaresSumWG calculates and prints sum of nums squared
// via sync.WaitGroup and sync.Mutex.
func SquaresSumMutex(nums []int) {
	sum := struct {
		value int64
		sync.Mutex
	}{}
	wg := sync.WaitGroup{}
	for _, n := range nums {
		n := int64(n)
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum.Lock()
			defer sum.Unlock()
			sum.value += n * n
		}()
	}
	wg.Wait()
	println("Sum:", sum.value)
}

// SquaresSumWG calculates and prints sum of nums squared
// via sync.WaitGroup and sync.Atomic.
func SquaresSumAtomic(nums []int) {
	var sum uint64
	wg := sync.WaitGroup{}
	for _, n := range nums {
		n := uint64(n)
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&sum, n*n)
		}()
	}
	wg.Wait()
	println("Sum:", sum)
}
