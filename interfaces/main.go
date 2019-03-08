package main

import (
	"fmt"
)

/*Defining Interfaces*/
type Guitarist interface {
	// PlayGuitar prints out "Playing Guitar"
	// to the terminal
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

/*Return Values*/
	// In real-world examples, we would typically have more complex functions within our interfaces that featured return values. 
	// In Go, we can define these interfaces like so:
type ReturnValues interface {
	Name() string
	Language() string
	Age() int
	Random() (string, error)
}

/*Satisfying Interfaces*/
type Employee interface {
	Language() string
	Age() int
	Random() (string, error)
}

type Engineer struct {
	Name string
}

func (e *Engineer) Language() string {
	return e.Name + " programs in Go."
}


func main()  {
	
/*
	Defining Interfaces
*/
fmt.Println("==============================================================Defining Interfaces:")
	var player BaseGuitarist
	player.Name = "Paul"
	player.PlayGuitar()

	var player2 AcousticGuitarist
	player2.Name = "Ringo"
	player2.PlayGuitar()
	
	// Should we wish, we could then create an array of type Guitarist which could store both our BaseGuitarist and AcousticGuitarist objects.
	var guitarists []Guitarist
	guitarists = append(guitarists, player)
	guitarists = append(guitarists, player2)
	fmt.Println("List Guitarist: ", guitarists)
	fmt.Println("Object 1 of Guitarist: ", guitarists[0])

/*
	Satisfying Interfaces
*/
fmt.Println("==============================================================Satisfying Interfaces:")
	// This will throw an error
	//	var programmers []Employee
	//	elliot := Engineer{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	//	programmers = append(programmers, elliot)
}