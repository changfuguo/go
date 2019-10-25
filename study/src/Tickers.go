package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done")
				return
			case t := <-timer.C:
				fmt.Println("Tick at ", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	timer.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
