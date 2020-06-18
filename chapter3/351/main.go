package main

import (
	"io"
	"os"
)

func main() {
	file, _ := os.Open("main.go")
	// reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(file, 1, 3)
	io.Copy(os.Stdout, sectionReader)
}
