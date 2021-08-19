package main

import (
	"errors"
	"fmt"
	"net/http"
)
const portNumber = ":8080"

// ----- HANDLER func ----- 
// In order for a function to respond to a request
// from a web browser, it has to handle 2 parameters,
// first- a response writer 
// second- has to handle a request

// Home - is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About - is the about page HANDLER  
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2,2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page, and 2 + 2 is %d", sum))
}

func Person(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Tanya Programmer")
}

// addValues - adds two integers together 
func addValues(x, y int) int {
	return x + y 
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return 
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil 
}


func main() {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/person", Person)
	http.HandleFunc("/divide", Divide)
	
	fmt.Println("starting application on port:", portNumber )
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber),"using fmt.Sprintln")
	// WEB SERVER that LISTENS for Requests in Go
	_ = http.ListenAndServe(portNumber, nil)

}