package main

import (
	"fmt"
	"strings"
)

func main() {

	testString := "Testing Strings"

	// Check a string contains a String
	fmt.Println(strings.Contains(testString, "st"))
	// Get index of value
	fmt.Println(strings.Index(testString, "es"))
	// Get the count of value within string
	fmt.Println(strings.Count(testString, "in"))
	/*
		Replace Values
		Params:
			- given string
			- value to replace
			- value to replace with
			- how many values to change
	*/
	fmt.Println(strings.Replace(testString, "s", "x", 1))
}
