package main

import (
	"fmt" 
	"net/http"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	//  n, err:= fmt.Fprintf(w, "Hello, Tanya Programmer")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes: %d",n))
	// })


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, Tanya Programmer, have a good day!")
	})

	// WEB SERVER that LISTENS for Requests in Go
	_ = http.ListenAndServe(":8080", nil)


















}