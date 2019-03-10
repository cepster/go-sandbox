package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var people = getSeed()

func main() {
	router := mux.NewRouter()

	/*
	 * MIDDLEWARE
	 */
	amw := MakeSampleAMW()
	router.Use(amw.Middleware)
	router.Use(loggingMiddleware)
	//////////////////////

	/*
	 * ROUTES
	 */
	router.HandleFunc("/people", getPeople)
	router.HandleFunc("/people/{id}", getPerson)
	//////////////////////

	//Start HTTP server
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		if item.idAsString() == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}
