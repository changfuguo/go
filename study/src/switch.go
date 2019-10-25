package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("here is ", i, "as")

	switch i {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	case 3:
		fmt.Println("case 3")
	}
	now := time.Now().Weekday()
	fmt.Println("now is ", now)
	switch now {
	case time.Saturday, time.Sunday:
		fmt.Println(" sleep")
	default:
		fmt.Println("work")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("is bool")
		case int:
			fmt.Println(" is int")
		case string:
			fmt.Println("is string")
		default:
			fmt.Println(" is default %T\n", t)
		}
	}
	whatAmI((123))
	whatAmI(true)
	whatAmI("hey")
	whatAmI(4.44)
}
