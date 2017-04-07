package main

import "fmt"

func main() {
	// If statements
	yourAge := 30

	if yourAge >= 18 {
		fmt.Println("You can drive")
	} else if yourAge >= 16 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't drive or vote")
	}
}
