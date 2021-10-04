package routes

import (
	admin "golang-rest-api/api/controllers/admin"
	bcrypt "golang-rest-api/api/controllers/bcrypt"
	jwt_token "golang-rest-api/api/controllers/jwtToken"

	"github.com/gorilla/mux"
)

var ApiRoutes = func(router *mux.Router) {

	//Admins
	router.HandleFunc("/getAdmins", admin.GetAdmins).Methods("GET")
	router.HandleFunc("/createAdmin", admin.CreateAdmin).Methods("POST")
	router.HandleFunc("/getAdminById/{adminId}", admin.GetAdminById).Methods("GET")
	router.HandleFunc("/deleteAdminById/{adminId}", admin.DeleteAdminById).Methods("DELETE")
	router.HandleFunc("/updateAdmin", admin.UpdateAdmin).Methods("PUT")

	//JWT Token
	router.HandleFunc("/createJWTToken", jwt_token.CreateJWTToken).Methods("POST")
	router.HandleFunc("/checkJWTToken", jwt_token.CheckJWTToken).Methods("POST")

	//Bcrypt
	router.HandleFunc("/cryptPassword", bcrypt.CryptPassword).Methods("POST")
	router.HandleFunc("/decryptPassword", bcrypt.DecryptPassword).Methods("POST")
}
