package main

import "fmt"

func main() {

	aSlice := []int{5, 4, 3, 2, 1}

	// Runs up to top number or ignores any not existing
	anotherSlice := aSlice[3:5]

	fmt.Println(anotherSlice[0])
	// If no inital val is given it is 0
	fmt.Println(aSlice[:3])

	// make( type, default values, max range)
	oneMoreSlice := make([]int, 5, 10)

	// copy another slice into this one
	copy(oneMoreSlice, aSlice)

	// We can also add values to the free space
	oneMoreSlice = append(oneMoreSlice, 0, -1)

	fmt.Println(oneMoreSlice[1])

}
