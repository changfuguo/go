package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

//我有一个key 还有一个管道
// 创建好之后往这管道扔值
type readOp struct {
	key  int
	resp chan int
}

// 除了key value 还有一个管道 写管道
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}

				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}

				writes <- write
				<-write.resp

				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

// 整个流程是这样d
// 1、write 写入到管道
// 2、写入之后读取设置state
// 3、write 完成
// 4、计数
// 同时有100个读读goroutiens
//5、读一个read
//6、从state 取出来key放入resp读管道
//7、从resp 管道读取
