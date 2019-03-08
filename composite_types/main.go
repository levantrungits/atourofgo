package main

import (
	"fmt"
)

/*Struct*/
type Person struct {
	name string
	age int
}
/*Nested Structs*/
	// our Boss struct
type Team struct {
	name string
	players [2]Person
}

func main() {

/*
	Arrays
*/
fmt.Println("==============================================================Arrays:")
	// declaring an empty array of strings
	var days []string
	fmt.Println(days)
	// declaring an array with elements
	daysElements := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(daysElements[0])
	fmt.Println(daysElements[5])

/*
	Slices
	- Slices in Go allow you to access a subset of an underlying array’s elements.
*/
fmt.Println("==============================================================Slices:")
	weekdays := daysElements[0:5]
	fmt.Println(weekdays) // This returns: [Monday Tuesday Wednesday Thursday Friday]

/*
	Maps
	- Maps are Go’s representation of hash tables
*/
fmt.Println("==============================================================Maps:")
	// For instance, let’s create a map of YouTube channel names to the numbers of subscribers for that channel:
	// This represents a mapping between a string data type and an int data type.
	youtubeSubscribers := map[string]int {
		"TutorialEdge": 	2240,
		"MKBHD": 			6580350,
		"Fun Fun Function": 171220,
	}
	fmt.Println(youtubeSubscribers["MKBHD"]) // prints out 6580350

/*
	Structs
	- In Go, we have this concept of a struct. These structs allow us to create data types that are aggregates of other data types.
*/
fmt.Println("==============================================================Structs:")
	// declaring a new `Person` ==> var myPerson Person
	// declaring a new `elliot`
	elliot := Person{name: "Elliot", age: 24}
	fmt.Println("Struct: ", elliot)
	// trying to roll back time to before I was injury prone
	elliot.age = 18
	fmt.Println("Struct: ", elliot)

/*
	Nested Structs
	- Structs are incredibly extensible due to the fact we can create nested structs within our structs. 
*/
fmt.Println("==============================================================Nested Structs:")
	// declaring an empty `Boss`
	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{
		Person{
			name: "Forrest",
		},
		Person{
			name: "Gordon",
		},
	}

	// declaring a boss with employees
	celtic := Team{
		name: "Celtic FC",
		players: players,
	}

	fmt.Println(celtic)








}
