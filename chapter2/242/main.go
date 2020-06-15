package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	os.Stdout.Write([]byte("hello"))
}
