package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	cpy, _ := os.Create("cpy.go")
	if err != nil {
		panic(err)
	}
	defer cpy.Close()
	defer file.Close()
	io.Copy(cpy, file)
}
