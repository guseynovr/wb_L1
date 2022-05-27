package main

import (
	"context"
	"time"
)

func main() {
	stopClose()
	stopPoll()
	stopContext()
}

func stopClose() {
	ch := make(chan struct{})
	finished := make(chan struct{})
	go func() {
		for _ = range ch {
			println("stopClose goroutine running")
		}
		println("stopClose goroutine finishing")
		finished <- struct{}{}
	}()

	time.Sleep(500 * time.Millisecond)
	for i := 0; i < 5; i++ {
		ch <- struct{}{}
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)  //signal goroutine to finish
	<-finished //wait for goroutine to finish
}

func stopPoll() {
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				println("stopPoll goroutine finishing")
				return
			default:
				println("stopPoll goroutine running")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	done <- struct{}{} //signal goroutine to finish
}

func stopContext() {
	ch := make(chan struct{})
	finished := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				println("stopContext goroutine finishing")
				finished <- struct{}{}
				return
			case <-ch:
				println("stopContext goroutine running")
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- struct{}{}
	}
	cancel()   //signal goroutine to finish
	<-finished //wait for goroutine to finish
}
