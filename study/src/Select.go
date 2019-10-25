package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("send one")

		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("send two")
		c2 <- "two"
	}()

	for i := 1; i <= 2; i++ {
		fmt.Println("i", i)
		select {
		case msg1 := <-c1:
			fmt.Println("recived", msg1)
		case msg2 := <-c2:
			fmt.Println("recvied", msg2)
		}
	}

	fmt.Println("game over")
}
