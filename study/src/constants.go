package main

import (
	"fmt"
	"math"
)

const s string = "changfuguo"

func main()  {
	fmt.Println(s)
	const n  = 50000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(math.Sin(d))
}
