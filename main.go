package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
	Port   int
}

func main() {
	// set application config
	var app application
	// read from command line

	// connect to the db

	app.Domain = "example.com"

	log.Printf("Starting server on %s:%d", app.Domain, port)

	// start the server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
