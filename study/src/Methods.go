package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

// 自动转化为指针，避免传value
func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}
func main() {
	r := rect{width: 12, height: 33}
	fmt.Println("area is", r.area())

	fmt.Println("perim is ", r.perim())

	rp := &r

	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
