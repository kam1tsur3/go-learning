package main

import (
	"io"
	"os"
)

func main() {
	io.CopyN(os.Stdout, os.Stdin, 8)

}
