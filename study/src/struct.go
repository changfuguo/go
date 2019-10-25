package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 111
	return &p
}
func main() {
	// 只能全部给 赋予值
	fmt.Println("person", Person{"Bob", 23})

	// 这种通过key  value 设置可以部分设置key
	fmt.Println("person no age", Person{name: "cfg", age: 33})
	fmt.Println("person no age", &Person{name: "cfg", age: 33})

	fmt.Println(newPerson("cfg"))

	p := Person{name: "cfg", age: 33}

	fmt.Println(p.name, p.age)

	sp := &p

	fmt.Println(sp.name)
}
