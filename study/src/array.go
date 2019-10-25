package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty ", a)

	a[4] = 100

	fmt.Println(a)
	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	fmt.Println("twoD", twoD)
	twoD[1][1] = 100
	fmt.Println("twoD", twoD)

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = (i + j) * i * j
			fmt.Println("i,j,value is", i, j, twoD[i][j])
		}
	}
}
