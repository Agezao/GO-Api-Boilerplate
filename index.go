package main

import "net/http"
import AppRouter "./routes"

func main() {
	// Binding routes
	AppRouter.Bind()

	// Starting server
	http.ListenAndServe(":8089", nil)
}
