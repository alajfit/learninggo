package main

import (
	"fmt"
)

// To have access to local variables
// External functions dont have access to local vars
func main() {
	num := 1

	doubleNum := func() int {
		num *= 2
		return num
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
}
