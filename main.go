// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Set up the routes
	router()

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
