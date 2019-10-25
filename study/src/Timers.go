package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println(timer1.C)
	<-timer1.C
	fmt.Println("timer1 expreid")

	timer2 := time.NewTimer(1 * time.Second)
	fmt.Println(timer2)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expreid")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println(stop2)
		fmt.Println("timer2 stop")
	}
}
