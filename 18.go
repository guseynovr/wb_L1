package main

import "sync"

type Counter struct {
	val int
	sync.RWMutex
}

func (c *Counter) increment() {
	c.Lock()
	defer c.Unlock()
	c.val++
}

func (c *Counter) value() int {
	c.RLock()
	defer c.RUnlock()
	return c.val
}

func main() {
	cnt := &Counter{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(c *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				c.increment()
			}
		}(cnt, wg)
	}

	wg.Wait()
	println("Result:", cnt.value())
}
