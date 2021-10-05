package routes

import (
	admin "golang-rest-api/api/controllers/admin"
	bcrypt "golang-rest-api/api/controllers/bcrypt"
	jwt_token "golang-rest-api/api/controllers/jwtToken"
	"golang-rest-api/api/controllers/mailer"
	rabbitmq_controlller "golang-rest-api/api/controllers/rabbitmq"
	"golang-rest-api/api/controllers/uuid"

	"github.com/gorilla/mux"
)

var ApiRoutes = func(router *mux.Router) {

	//Admins
	router.HandleFunc("/getAdmins", admin.GetAdmins).Methods("GET")
	router.HandleFunc("/createAdmin", admin.CreateAdmin).Methods("POST")
	router.HandleFunc("/getAdminById/{adminId}", admin.GetAdminById).Methods("GET")
	router.HandleFunc("/deleteAdminById/{adminId}", admin.DeleteAdminById).Methods("DELETE")
	router.HandleFunc("/updateAdmin", admin.UpdateAdmin).Methods("PUT")

	//RabbitMQ
	router.HandleFunc("/rabbitmqSendDataToQueue", rabbitmq_controlller.RabbitmqSendDataToQueue).Methods("POST")
	router.HandleFunc("/receiveDatasFromRabbitMQ", rabbitmq_controlller.ReceiveDatasFromRabbitMQ).Methods("POST")

	//JWT Token
	router.HandleFunc("/createJWTToken", jwt_token.CreateJWTToken).Methods("POST")
	router.HandleFunc("/checkJWTToken", jwt_token.CheckJWTToken).Methods("POST")

	//Bcrypt
	router.HandleFunc("/cryptPassword", bcrypt.CryptPassword).Methods("POST")
	router.HandleFunc("/decryptPassword", bcrypt.DecryptPassword).Methods("POST")

	//Mailer
	router.HandleFunc("/sendMail", mailer.SendMail).Methods("POST")

	//uuid
	router.HandleFunc("/generateUUID", uuid.GenerateUUID).Methods("GET")
}
