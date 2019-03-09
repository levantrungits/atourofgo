package main

import (
	"strconv"
	"io/ioutil"
	"fmt"
	"os"
	// import our encoding/json package
	"encoding/json"
)

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter string 	`json:"twitter"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age int 	`json:age`
	Social Social `json:"social"`
}

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

/*Reading and Parsing a JSON File*/
func main() {

/*
	Reading and Parsing a JSON File
*/
fmt.Println("==============================================================Reading and Parsing a JSON File:")
	// Open our jsonFile
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error the handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

/*
	Parsing with Structs
*/
fmt.Println("==============================================================Unmarshalling our JSON:")
	// read our opened xmlFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)
	
	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i:=0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

/*
	Working with Unstructured Data
		use standard interfaces{} in order to read in any JSON data
*/
fmt.Println("==============================================================Working with Unstructured Data:")
	// Open our jsonFIle2
	jsonFile2, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile2.Close()

	byteValue2, _ := ioutil.ReadAll(jsonFile2)
	
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue2), &result)
	fmt.Println(result["users"])
	
}