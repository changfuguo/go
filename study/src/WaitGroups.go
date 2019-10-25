package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf(" worker %d starting \n", id)

	time.Sleep(time.Second)

	fmt.Printf("worker %d done\n", id)

	// 通知group 当前的wait 结束了
	//wg.Done()
}

func main() {
	//WaitGroup 只能通过指针传递
	//使用这个WaitGroup 等到所有的从这里启动的goroutines 结束
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		// 增加一个routines并且增加WaitGroup的计数器
		wg.Add(1)
		go worker(i, &wg)
	}

	//hang住直到所有的WaitGroup计数器变为0；所有的worker 通知他们自己已经完成
	wg.Wait()

}
