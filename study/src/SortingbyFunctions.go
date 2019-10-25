package main

import (
	"fmt"
	"sort"
)

type bytesLength []string

//对特定类型读数据排序，必须实现
// sort.Interface - Len, Less, and Swap
// 只有这三种方法都实现了，才认为你实现了一个可被排序的数据接口

// 这里都是引用传递
func (b bytesLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b bytesLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}
func (b bytesLength) Len() int {
	return len(b)
}
func main() {
	fruits := []string{"yellow", "bananan", "apple"}
	sort.Sort(bytesLength(fruits))
	bytesLength(fruits).Swap(1, 2)
	fmt.Println(fruits)
}
