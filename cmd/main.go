package main

import (
	"fmt"
	"net/http"

	"github.com/William9923/microservice-server/handler"
)

func main() {

	handler.Route()

	// Start the HTTP server on port 8080
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
