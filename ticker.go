package main

import "time"

// The time between each tick
var interval = time.Second / 10

// Send a tick counter into the channel
func runTicker(ch chan int) {
	i := 0
	for _ = range time.Tick(time.Second / 10) {
		i++
		ch <- i
	}
}

func newTicker() chan int {
	ch := make(chan int, 10)
	go runTicker(ch)
	return ch
}
