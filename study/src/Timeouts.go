package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		//c1 <- "result 1"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("from c1", msg1)
	case <-time.After(time.Second * 1):
		fmt.Println("default a after")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(4 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
