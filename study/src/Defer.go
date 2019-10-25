package main

import (
	"fmt"
	"os"
)

//
func createFile(p string) *os.File {
	fmt.Println("createing file")
	f, err := os.Create(p)

	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprint(f, "hi i am data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprint(os.Stderr, "error %v\n", err)
		os.Exit(1)
	}
}
func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
