package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/rabbitmq"
	"net/http"
)

func Routes() {
	rabbit := rabbitmq.RabbitMQConnect()

	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/tokenControl", controllers.TokenControl)
	http.HandleFunc("/bcryptPassword", controllers.BcryptPassword)
	http.HandleFunc("/decryptPassword", controllers.DecryptPassword)
	http.HandleFunc("/rabbitmqSendDataToQueue", rabbit.RabbitmqSendDataToQueue)
	http.HandleFunc("/rabbitmqReceiveDataFromQueue", rabbit.ReceiveDatasFromRabbitMQ)
}
