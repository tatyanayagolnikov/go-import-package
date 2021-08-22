package main

import (
	"fmt"
	"net/http"
	"github.com/tatyanayagolnikov/go-website/pkg/handlers"
)
const portNumber = ":8080"




func main() {
	
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/person", handlers.Person)
	
	
	fmt.Println("starting application on port:", portNumber)
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber),"using fmt.Sprintln")
	// WEB SERVER that LISTENS for Requests in Go
	_ = http.ListenAndServe(portNumber, nil)

	

}