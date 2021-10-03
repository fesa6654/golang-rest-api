package routes

import (
	"golang-rest-api/api/controllers/admin"

	"github.com/gorilla/mux"
)

var ApiRoutes = func(router *mux.Router) {
	router.HandleFunc("/getAdmins", admin.GetAdmins).Methods("GET")
	router.HandleFunc("/createAdmin", admin.CreateAdmin).Methods("POST")
	router.HandleFunc("/getAdminById/{adminId}", admin.GetAdminById).Methods("GET")
	router.HandleFunc("/deleteAdminById/{adminId}", admin.DeleteAdminById).Methods("DELETE")
	router.HandleFunc("/updateAdmin", admin.UpdateAdmin).Methods("PUT")
}
