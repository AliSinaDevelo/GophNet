package main

import (
	"fmt"
	"net/http"
	"time"
	"GophNet/server"
)

func main() {
	port := 8080
	// create the server
	server.SetupRoutes()

	// run the server
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting GophNet server on http://localhost:%d\n", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

