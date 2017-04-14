package main

import "fmt"

func main() {
	fmt.Println(safeDiv(1, 0))
	fmt.Println(safeDiv(6, 2))
	callPanic()
}

func safeDiv(num1, num2 int) int {

	// NB: Recover allows our program to finish
	// 		even if a fatal error should occur
	defer func() {
		// This line will print the error
		//fmt.Println(recover())

		// This line will pass right over the error
		recover()
	}()

	solution := num1 / num2
	return solution
}

func callPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	// panic is used to force an error
	panic("PANIC")
}
