package main

import "fmt"

func main() {
	// For Loop
	i := 1

	for i <= 10 {
		fmt.Println(i)

		i++
	}

	// Another way to declare a for Loop
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
}
