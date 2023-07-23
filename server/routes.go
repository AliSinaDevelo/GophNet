package server

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Howdy, GophNet!")
}

// aboutHandler handles requests to the "/about" endpoint and responds with an about message
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GophNet! This is the about page!")
}