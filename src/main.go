package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ryanvoortman/go-rest-api/src/app"
	database2 "ryanvoortman/go-rest-api/src/database"
)

func main() {
	database, err := database2.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe("localhost:8080", app.Router))
}
