package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// Splitting Comma delimmeted Strings
	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))

	// Sorting list of letters
	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters: ", listOfLetters)

	// Join a List by a value
	listOfNums := strings.Join([]string{"3", "2", "1"}, ",")
	fmt.Println(listOfNums)
	listOfOtherNums := strings.Join([]string{"3", "2", "1"}, "")
	fmt.Println(listOfOtherNums)
}
