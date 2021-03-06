package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("User ID: %d\n", os.Getuid())
	fmt.Printf("Group ID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("sub group ID: %v\n", groups)
}
