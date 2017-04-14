package main

import "fmt"

func main() {
	// Defer to hold off on calling a function
	// 	Till the enclosing function finishes
	//		e.g. main
	defer printTwo()
	printOne()
}

func printOne() {
	fmt.Println(1)
}
func printTwo() {
	fmt.Println(2)
}
