package routes

import (
	"go-rest-api/controllers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/tokenControl", controllers.TokenControl)
}
