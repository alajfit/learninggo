package main

import "fmt"

func main() {
	// Strings use double quotes or back tick quotes
	var myName string = "Alan Fitzpatrick"
	fmt.Println(myName, "length is", len(myName))

	// String concatonation is done with a plus sign
	fmt.Println(myName + ` is a Developer`)
}
