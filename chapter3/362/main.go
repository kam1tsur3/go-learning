package main

import (
	"fmt"
	"strings"
)

var source = "123 1.234 1.0e4 test"

func main() {
	var d int
	var f32 float32
	var f64 float64
	var s string
	reader := strings.NewReader(source)
	fmt.Fscan(reader, &d, &f32, &f64, &s)
	fmt.Printf("%v, %v, %v, %v", d, f32, f64, s)

}
