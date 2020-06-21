package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {

	var stream io.Reader
	/*
		m_reader := io.MultiReader(computer, system, programming)
		r,w := io.Pipe()
		go pass := io.MultiReader(m_reader, r)
		teeReader := io.TeeReader(pass, w)
		for {
			buffer := []byte
			go pass.Read()
		}
	*/
	a := io.NewSectionReader(programming, 5, 1)
	s := io.LimitReader(system, 1)
	c := io.LimitReader(computer, 1)
	i1 := io.NewSectionReader(programming, 8, 1)
	i2 := io.NewSectionReader(programming, 8, 1)
	stream = io.MultiReader(a, s, c, i1, i2)
	io.Copy(os.Stdout, stream)
}
