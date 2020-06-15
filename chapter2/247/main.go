package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with at %v", time.Now())
}
