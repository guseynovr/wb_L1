package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Time is not specified.")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		log.Fatal("Specify time (>0) in seconds.")
 	}

	ch := make(chan int)
	stop := time.After(time.Duration(n) * time.Second)
	go func() {
		for i := range ch {
			println("received ", i)
		}
	}()

	i := 0
	for {
		select {
		case <-stop:
			close(ch)
			return
		case ch <- i:
			time.Sleep(500 * time.Millisecond)
			i++	
		}
	}
}
