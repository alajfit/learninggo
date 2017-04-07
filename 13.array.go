package main

import (
	"fmt"
)

func main() {
	var favNums [2]float64

	favNums[0] = 163
	favNums[1] = 1.2345

	fmt.Println(favNums[0])

	favNames := [2]string{"Natasha", "Alan"}

	for i, value := range favNames {
		fmt.Println(value, i)
	}

	// Use underscore to ignore the index
	for _, value := range favNames {
		fmt.Println(value)
	}
}
