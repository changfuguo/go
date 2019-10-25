package main

//check
import (
	"fmt"
	"time"
	//"time"
)

/*
这里会演示一些worker的并发实例，，这些workers 会接受一些工作任务通道，并且发送相应的结果到results
为了模拟消耗资源的昂贵任务，在每一项工作时休眠一秒钟
*/

func worker(wid int, jobs <-chan int, result chan<- int) {
	fmt.Println("in worker")
	for j := range jobs {
		fmt.Println("worker ", wid, "start job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", wid, "finished", j)

		result <- j * 2
	}
}

func main() {
	// 为了使用workers 池，我们给他们发送任务并搜集结果
	// 这里使用了两个chan
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 有三个工人开工，并被block，因为没有job
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 有五项工作要做
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
		jobs <- j
	}
	close(jobs)
	// 最后的结果，

	for i := 1; i <= 5; i++ {
		<-results
	}
}
