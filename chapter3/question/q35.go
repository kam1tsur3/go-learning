package main

import (
	"io"
	"os"
	"bytes"
)

func myCopyN(dest io.Writer, src io.Reader, length int) {
	buffer := make([]byte, 20)
	off := 0
	for {
		read_n, _ := src.Read(buffer)
		if (off + read_n) < length {
			rw_buf := bytes.NewBuffer(buffer)
			io.Copy(dest, rw_buf)
		} else {
			rw_buf := bytes.NewBuffer(buffer[:(length-off)])
			io.Copy(dest, rw_buf)
			break
		}
		off += read_n	
		if read_n < len(buffer) {
			break
		}
	}
}

func main() {
	file, _ := os.Open("q35.go")
	myCopyN(os.Stdout, file, 30)
}
