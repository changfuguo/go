package main

import "fmt"

func zeroValue(n int) {
	n = 0
}

func zeroPtr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroValue(i)

	fmt.Println("zeroValue:", i)

	zeroPtr(&i)

	fmt.Println("zeroPtr:", i)

	fmt.Println("zeroPtr &:", &i)
}
