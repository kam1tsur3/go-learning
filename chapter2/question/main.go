package main

import (
	"fmt"
	"os"
)

func main() {
	// Q 2.1
	a := 3
	b := "KM2"
	file, err := os.Create("Q2.1.txte")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%s, %d", b, a)
	fmt.Fprintf(os.Stdout, "%s, %d", b, a)
}
