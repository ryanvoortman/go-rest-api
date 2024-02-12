package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setUpRouter(router *mux.Router) {
	router.Methods("Post").Path("/endpoint").HandlerFunc(postFunction)
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("You called a thing")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	setUpRouter(router)

	log.Fatal(http.ListenAndServe("8080", router))

}
