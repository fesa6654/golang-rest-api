package main

import (
	"golang-rest-api/api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.ApiRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:5000", r))
}
