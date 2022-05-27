package main

import (
	"os"
	"strings"
	"sync"
)

type safeMap struct {
	count map[rune]int
	sync.Mutex
}

func main() {
	// raw, err := os.ReadFile("WB Tech_ level # 1 (Golang).txt")
	raw, err := os.ReadFile("test")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(raw), "\n")

	wg := &sync.WaitGroup{}
	sm := &safeMap{count: make(map[rune]int)}
	for _, line := range lines {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			for _, r := range line {
				sm.Lock()
				sm.count[r] = sm.count[r] + 1
				sm.Unlock()
			}
		}(line)
	}
	wg.Wait()
	for key, value := range sm.count {
		println(string(key), value)
	}
}
