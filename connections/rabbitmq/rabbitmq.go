package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

var (
	RabbitConnect *amqp.Connection
	RabbitChannel *amqp.Channel
)

func RabbitMQConnect() {

	con, err := amqp.Dial("amqps://oryeecvq:Ba5zvf5J9JD5tObauu054nnssDfEJPXh@cow.rmq2.cloudamqp.com/oryeecvq")

	if err != nil {
		fmt.Printf("RabbitMQ Connection is not Running !\n")
	} else {
		fmt.Printf("RabbitMQ Connection is Running !\n")
	}

	RabbitConnect = con

	ch, err := con.Channel()

	if err != nil {
		fmt.Println("RabbitMQ Channel Error !")
	}

	RabbitChannel = ch

}
