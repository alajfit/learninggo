package main

import "fmt"

func main() {

	// Pass key value type
	presAge := make(map[string]int)
	presAge["TheodoreRoosevelt"] = 42

	fmt.Println(presAge["TheodoreRoosevelt"])

	// Get The Length of a Map
	mapLen := len(presAge)
	fmt.Println(mapLen)

	// remove element from a Map
	delete(presAge, "TheodoreRoosevelt")
	fmt.Println(len(presAge))
}
