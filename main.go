package main

import (
	appRouter "./routes"
	"log"
	"net/http"
)

func main() {
	// Binding routes
	appRouter.Bind()

	// Starting server
	log.Fatal(http.ListenAndServe(":8089", nil))
}
