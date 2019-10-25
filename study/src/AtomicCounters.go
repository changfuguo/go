package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
在atomic包中对几种基础类型提供了原子操作，
包括int32，int64，uint32，uint64，uintptr，unsafe.Pointer。
对于每一种类型，提供了五类原子操作分别是

Add, 增加和减少
CompareAndSwap, 比较并交换
Swap, 交换
Load , 读取
Store, 存储
*/
func main() {

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(ops)
}
