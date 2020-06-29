package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.After(10 * time.Second)
	fmt.Println("timer started")
	<-timer
	fmt.Println("timer ended")
}
