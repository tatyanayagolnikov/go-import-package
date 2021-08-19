package main

import (
	"fmt"
	"net/http"
	"text/template"
)
const portNumber = ":8080"

// ----- HANDLER func ----- 
// In order for a function to respond to a request
// from a web browser, it has to handle 2 parameters,
// first- a response writer 
// second- has to handle a request

// Home - is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl.html")
}

// About - is the about page HANDLER  
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl.html")
}

func Person(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "person.page.tmpl.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return 
	}
}


func main() {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/person", Person)
	
	
	fmt.Println("starting application on port:", portNumber )
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber),"using fmt.Sprintln")
	// WEB SERVER that LISTENS for Requests in Go
	_ = http.ListenAndServe(portNumber, nil)

}