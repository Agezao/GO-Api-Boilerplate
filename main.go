package main

import (
	"log"
	"net/http"
	"strconv"

	. "./config"
	ctx "./contexts/mongodb"
	appRouter "./routes"
)

func main() {
	//Connecting to db
	ctx.Connect()

	// Binding routes
	appRouter.Bind()

	// Starting server
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil))
}
