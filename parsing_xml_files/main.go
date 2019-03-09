package main

import (
	"io/ioutil"
	"os"
	"encoding/xml"
	"fmt"
)

// a simple struct which contains all our
// social links
type Social struct {
	XMLName xml.Name 	`xml:"social"`
	Facebook string 	`xml:"facebook"`
	Twitter string 		`xml:"twitter"`
	Youtube string		`xml:"youtube"`
}

// the user struct, thi contains our
// Type attribute, our user's name and
// a social struct which will contain all our social links
type User struct {
	XMLName xml.Name	`xml:"user"`
	Type string			`xml:"type,attr"`
	Name string			`xml:"name"`
	Social Social		`xml:"social"`
}

// our struct which contains the complete
// array of all Users in the file
type Users struct {
	XMLName xml.Name	`xml:"users"`
	Users []User		`xml:"user"`
}

func main()  {

/*
	Working with Unstructured Data
		use standard interfaces{} in order to read in any JSON data
*/
fmt.Println("==============================================================Working with Unstructured Data:")
	// Open our xmlFile
	xmlFile, err := os.Open("users.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.xml")
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// inittialize our Users array
	var users Users

	// unmarshal our byteArray which contains our
	xml.Unmarshal(byteValue, &users)

	// iterate through every user
	for i:=0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}


}