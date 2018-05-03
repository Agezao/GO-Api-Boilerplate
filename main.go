package main

import (
	. "./config"
	appRouter "./routes"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Binding routes
	appRouter.Bind()

	// Starting server
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil))
}
