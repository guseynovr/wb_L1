package main

import "sync"

func main() {
	nums := []int{2, 4, 6, 8, 10}
	printSquares(nums)
	// printSquaresWG(nums)
}

// printSquares prints squares of every number of the given
// nums slice concurrently, using an unbuffered channel.
func printSquares(nums []int) {
	squares := make(chan int)
	for _, n := range nums {
		go func(n int) {
			squares <- n * n
		}(n)
	}
	for i := 0; i < len(nums); i++ {
		println(<-squares)
	}
}

// printSquaresWG prints squares of every num in nums
// concurrently using sync.WaitGroup.
func printSquaresWG(nums []int) {
	wg := &sync.WaitGroup{}
	for _, n := range nums {
		wg.Add(1)
		go func (n int, wg *sync.WaitGroup) {
			defer wg.Done()
			println(n * n)
		}(n, wg)
	}
	wg.Wait()
}
