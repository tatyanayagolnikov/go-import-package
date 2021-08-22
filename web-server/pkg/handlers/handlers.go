package handlers

import (
	"net/http"
	"github.com/tatyanayagolnikov/go-website/pkg/render"
	
)	


// ----- HANDLER func ----- 
// In order for a function to respond to a request
// from a web browser, it has to handle 2 parameters,
// first- a response writer 
// second- has to handle a request

// Home - is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl.html")
}

// About - is the about page HANDLER  
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl.html")
}

// Person - is the person page handler
func Person(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "person.page.tmpl.html")
}
