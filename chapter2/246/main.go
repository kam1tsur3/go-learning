package main

import (
	"bufio"
	"os"
)

func main() {
	/*
		file, err := os.Create("multiwriter.txt")
		if err != nil {
			panic(err)
		}
		writer := io.MultiWriter(file, os.Stdout)
		io.WriteString(writer, "io.MultiWrite example\n")
	*/
	/*
		file, err := os.Create("test.txt.gz")
		if err != nil {
			panic(err)
		}
		writer := gzip.NewWriter(file)
		writer.Header.Name = "test.txt"
		io.WriteString(writer, "gzip.Writer example\n")
		writer.Close()
	*/
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	//buffer.Flush()
}
