package main

import "fmt"

func main() {
	yourAge := 30

	switch yourAge {
	case 29:
		fmt.Println("Nearly Middle Aged")
	case 30:
		fmt.Println("Middle Aged")
	case 31:
		fmt.Println("Over Middle Aged")
	default:
		fmt.Println("Not an important Age")
	}
}
