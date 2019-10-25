package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("nums ,sum:", nums, sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("here is num 3", i)
		}
	}

	kvs := map[string]string{"name": "changfuguo", "address": "china"}
	for k, v := range kvs {
		fmt.Println("k:,", k, ". v", v)
	}

	for k := range kvs {
		fmt.Println('k', k)
	}

	for _, v := range kvs {
		fmt.Println("v", v)
	}

	for _, v := range "changfuguo" {
		fmt.Println(v)
	}
}
