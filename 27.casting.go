package main

import (
	"fmt"
	"strconv"
)

func main() {

	randInt := 5
	randFloat := 10.5
	randString := "100"

	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))

	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)
}
