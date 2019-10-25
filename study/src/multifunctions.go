package main

import "fmt"

func vals() (int, int, int) {
	return 3, 7, 7
}

func main() {

	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	_, d, c := vals()
	fmt.Println(c, d)
}
