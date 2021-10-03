package main

import (
	"golang-rest-api/api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()

	routes.ApiRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:5000", r))
}
