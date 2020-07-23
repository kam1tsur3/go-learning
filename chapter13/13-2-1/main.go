package main

import (
	"fmt"
	"time"
)

func subl(c int) {
	fmt.Println("share by arguments:",  c*c)
}

func main() {
	go subl(10)

	c:= 20
	go func() {
		fmt.Println("share by capture", c*c)
	}()
	time.Sleep(time.Second)
}
