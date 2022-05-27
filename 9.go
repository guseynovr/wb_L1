package main

func main() {
	sl := make([]int, 20)
	for i := range sl {
		sl[i] = i
	}
	nums := make(chan int)
	doubled := make(chan int)

	go putNums(sl, nums)
	go doubleNums(nums, doubled)

	for d := range doubled {
		println(d)
	}
}

// putNums reads numbers from sl and writes them to nums.
// nums channel is closed when all numbers ar written.
func putNums(sl []int, nums chan<- int) {
	for _, n := range sl {
		nums <- n * 2
	}
	close(nums)
}

// doubleNums reads numbers from nums, multiplies them by 2 and writes
// them to doubled. nums channel is closed when all numbers ar written.
func doubleNums(nums <-chan int, doubled chan<- int) {
	for n := range nums {
		doubled <- n
	}
	close(doubled)
}
