package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// j一直接受chan 传过来的值，如果没有close more为true ，关闭了就是false
			// 这个时候再结束done
			j, more := <-jobs
			fmt.Println("more %d", more)
			if more {
				fmt.Println("recived jobs", j)
			} else {
				fmt.Println("recived all job")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	<-done
}
