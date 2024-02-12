package app

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ryanvoortman/go-rest-api/src/database"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Path("/endpoint/{id}").
		HandlerFunc(app.getFunction)
	app.Router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(app.postFunction)
}

func (app *App) postFunction(writer http.ResponseWriter, request *http.Request) {
	_, err := app.Database.Exec("INSERT INTO `test` (name) VALUES ('myname')")
	if err != nil {
		log.Fatal("Database insert failed")
	}

	log.Println("You called a thing")
	writer.WriteHeader(http.StatusOK)
}

func (app *App) getFunction(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No id in path")
	}

	dbdata := &database.DbData{}
	err := app.Database.
		QueryRow("SELECT id, date, name FROM `test` WHERE id = ?", id).
		Scan(&dbdata.ID, &dbdata.Date, &dbdata.Name)
	if err != nil {
		log.Fatal("Database select failed")
	}

	log.Println("you fetched a thing")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(dbdata); err != nil {
		panic(err)
	}
}
