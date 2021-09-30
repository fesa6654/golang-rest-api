package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/mysql"
	"go-rest-api/rabbitmq"
	"net/http"
)

func Routes() {

	var myCon mysql.MySQL
	var rabbitCon rabbitmq.Rabbit

	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/tokenControl", controllers.TokenControl)
	http.HandleFunc("/bcryptPassword", controllers.BcryptPassword)
	http.HandleFunc("/decryptPassword", controllers.DecryptPassword)
	http.HandleFunc("/rabbitmqSendDataToQueue", rabbitCon.RabbitmqSendDataToQueue)
	http.HandleFunc("/rabbitmqReceiveDataFromQueue", rabbitCon.ReceiveDatasFromRabbitMQ)
	http.HandleFunc("/getData", myCon.GetData)

}
