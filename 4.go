package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Number of workers is not specified.")
	}
	poolSize, err := strconv.Atoi(os.Args[1])
	if err != nil || poolSize < 1 {
		log.Fatal("Invalid argument for number of workers. Must be > 0.")
	}
	input := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go startWorker(input, i + 1, wg)
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	i := 0

Loop:
	for {
		select {
		case <-stop:
			close(input)
			break Loop
		case input <- i:
			i++
		}
	}

	wg.Wait()
	fmt.Println("Received SIGINT: closing input. Last data:", i - 1)
}

func startWorker(input <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		fmt.Printf("worker: %d, data: %d\n", id, data)
		// time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf("worker %d finished\n", id)
}
