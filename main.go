package main

import (
	AppRouter "./routes"
	"log"
	"net/http"
)

func main() {
	// Binding routes
	AppRouter.Bind()

	// Starting server
	log.Fatal(http.ListenAndServe(":8089", nil))
}
