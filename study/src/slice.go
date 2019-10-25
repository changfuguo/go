package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "c"
	s[1] = "f"
	s[2] = "g"
	fmt.Println(s)
	fmt.Println(len(s))

	s = append(s, "lyl", "cjy")
	fmt.Println(s)
	fmt.Println(len(s))

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)

	fmt.Println(c[:2])
	fmt.Println(c[2:3])
	fmt.Println(c[2:])

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
