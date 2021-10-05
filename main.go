package main

import (
	"golang-rest-api/api/routes"
	"golang-rest-api/connections/rabbitmq"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	rabbitmq.RabbitMQConnect()

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	routes.ApiRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:5000", handler))
}
