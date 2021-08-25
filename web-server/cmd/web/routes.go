package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/tatyanayagolnikov/go-website/pkg/config"
	"github.com/tatyanayagolnikov/go-website/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()  // multiplexer or http handler

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/person", http.HandlerFunc(handlers.Repo.Person))

	return mux
}

