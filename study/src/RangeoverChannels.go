package main

import "fmt"

func main() {

	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	//<-queue
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
