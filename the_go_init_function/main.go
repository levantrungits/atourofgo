package main

import (
	"fmt"
)

var name string

/*Awesome, so with this working, we can start to do cool things such as variable initialization.*/
func init() {
	fmt.Println("This will get called on main initialization")
	name = "levantrungits"
}

/*Init Function: it is only called once.*/
func main1()  {
/*
	The Init Function
*/
fmt.Println("==============================================================The Init Function:")
	fmt.Println("My Wonderful Go Program")
	fmt.Printf("Name: %s\n", name)
}