package main

import "time"

func main() {

	println("Starting timer for 2 seconds, time:", time.Now().String())
	mSleep(2 * time.Second)
	// mSleep2(2 * time.Second)
	println("Timer stopped, time:", time.Now().String())
}

func mSleep(d time.Duration) {
	<-time.After(d)
}

func mSleep2(d time.Duration) {
	ticker := time.NewTicker(d)
	<-ticker.C
}
