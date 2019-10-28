package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	//可以预编译
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	//[0 5]
	fmt.Println(r.FindStringIndex("peach punch"))

	//[peach ea]
	fmt.Println(r.FindStringSubmatch("peach punch"))
	//[0 5 1 3]
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	//[peach punch pinch]
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	//[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	//[peach punch]
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
