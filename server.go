package main

import (
	"fmt"
	"net/http"

	"go-rest-api/routes"
)

func startServer() {
	port := 5000
	fmt.Printf("HTTP API is Running on %v Port !\n", port)
	http.ListenAndServe(":5000", nil)
}

func main() {
	routes.Routes()
	startServer()
}
