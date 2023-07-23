package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	port := 8080
	// create the server
	http.HandleFunc("/", helloHandler)

	// run the server
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting GophNet server on http://localhost:%d\n", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

// helloHandler handles requests to the root URL
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Howdy, GophNet!")
}