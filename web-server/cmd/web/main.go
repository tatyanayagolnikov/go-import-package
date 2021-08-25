package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tatyanayagolnikov/go-website/pkg/config"
	"github.com/tatyanayagolnikov/go-website/pkg/handlers"
	"github.com/tatyanayagolnikov/go-website/pkg/render"
)
const portNumber = ":8080"




func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false 

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)


	
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/person", handlers.Repo.Person)
	
	
	fmt.Println("starting application on port:", portNumber)
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber),"using fmt.Sprintln")
	// WEB SERVER that LISTENS for Requests in Go
	_ = http.ListenAndServe(portNumber, nil)

	

}