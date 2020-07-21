package main

import (
	"fmt"
	"os"
)

func main() {
	path, _ := os.Executable()
	fmt.Printf("file name: %s\n", os.Args[0])
	fmt.Printf("file path: %s\n", path)
}
