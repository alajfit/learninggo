package main

import "fmt"

func main() {
	listNums := []int{1, 5, 6}
	fmt.Println(sumOfNums(listNums))
	fmt.Println(next2Values(100))
	fmt.Println(getAverage(2, 5, 7, 10))
}

// Functions in Go
func sumOfNums(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

// Multiple return values
func next2Values(number int) (int, int) {
	return number + 1, number + 2
}

// Unknown amount of args
func getAverage(args ...int) int {
	var finalValue int
	var runner int

	for _, val := range args {
		finalValue += val
		runner++
	}

	finalValue = finalValue / runner
	return finalValue
}
