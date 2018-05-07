package main

import (
	"log"
	"net/http"
	"strconv"

	. "./config"
	appRouter "./routes"
)

func main() {
	// Binding routes
	appRouter.Bind()

	// Starting server
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil))
}
