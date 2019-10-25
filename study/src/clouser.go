package main

import (
	"fmt"
)

// 这里的这里的y意思返回一个函数函数，所以yao指定指定返回函数的类
func intSeq() func() int {
	i := 1
	return func() int {
		i++
		return i
	}
}
func main() {
	nextId := intSeq()

	fmt.Println(nextId())
	fmt.Println(nextId())
	fmt.Println(nextId())

	newInts := intSeq()
	fmt.Println(newInts())
}
