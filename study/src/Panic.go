package main

import "os"

// 通常用于一个返回不知道要如何处理的错误
func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
