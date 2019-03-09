package main

import (
	"fmt"
)

// this variable is initialized first due to
// order of declaration
var initCounter int

func init() {
	fmt.Println("Called First in Order of Declaration")
	initCounter++
}

func init() {
	fmt.Println("Called second in order of declaration")
	initCounter++
}

/*
	Go does indeed support having 2 separate init() functions within the one file.
	The one caveat of this style is that you will have to take care when ensuring declaration order.
*/
func main() {
	fmt.Println("Does nothing of any significance")
	fmt.Printf("Init Counter: %d\n", initCounter)
}