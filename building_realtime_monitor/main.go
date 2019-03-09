/*
	Creating a real-time YouTube stats monitoring system in Go
	Creating a WebSocket server and using WebSockets to communicate in real-time with a frontend application, 
	as well as how you can interact with an existing REST API to get the subscriber stats we need.
	TODO----
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

// homePage will be a simple "hello World" style page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// setupRoutes handles setting up our servers
// routes and matching them to their respective
// functions
func setupRoutes() {
	http.HandleFunc("/", homePage)
	// here we kick off our server on localhost:8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// our main function
func main()  {
	fmt.Println("Youtube Subscriber Monitor")

	// calls setup routes
	setupRoutes()
}