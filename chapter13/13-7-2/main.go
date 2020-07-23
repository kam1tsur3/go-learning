package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Job 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Job 2")
		wg.Done()
	}()
	
	wg.Wait()
	fmt.Println("finish")
}
