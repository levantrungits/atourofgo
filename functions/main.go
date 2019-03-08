package main

import (
	"fmt"
)

/*Function Declaration*/
func myfunction(firstName string, lastName string) (string, error) {
	return firstName + " " + lastName, nil
}

/*Anonymous Functions*/
func addOne() func() int {
	var x int
	return func() int {
		x++
		return x+1
	}
}

func main()  {
	
/*
	Function Declaration
	func name(parameter-list) (result-list) {
		body
	}
*/
fmt.Println("==============================================================Function Declaration:")
	fullName, err := myfunction("Elliot", "Forbes")
	if err != nil {
		fmt.Println("Handle Error Case.")
	}
	fmt.Println("My name is ", fullName)

/*
	Anonymous Functions
	Anonymous functions are very similar to regular functions except they lack a named identifier. 
	These functions can be defined within named functions and can have access to any variables within it’s enclosing function like so:
*/
fmt.Println("==============================================================Anonymous Functions:")
	myFunc := addOne()
	fmt.Println(myFunc()) // 2
	fmt.Println(myFunc()) // 3
	fmt.Println(myFunc()) // 4
	fmt.Println(myFunc()) // 5
}