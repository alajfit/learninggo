package main

import "fmt"

// Pointers allow us to send a pointer to a memory address

func main() {
	x := 0
	yPtr := new(int)

	// We can send the memory ref with &
	changeXVal(&x)
	changeYVal(yPtr)

	fmt.Println(x)
	fmt.Println(*yPtr)

	// We can actually print the memory address
	fmt.Println(&x)
}

// We can tell the func its getting a memory address with *
func changeXVal(x *int) {
	// and then use the address with *
	*x = 2
}

func changeYVal(yPtr *int) {
	*yPtr = 100
}
