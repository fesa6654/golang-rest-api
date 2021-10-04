package rabbitmq_controller

import (
	"fmt"
	"golang-rest-api/connections/rabbitmq"
	"net/http"

	"github.com/streadway/amqp"
)

var channel = rabbitmq.RabbitMQChannel()

func RabbitmqSendDataToQueue(w http.ResponseWriter, r *http.Request) {

	_, err := channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
	}

	// attempt to publish a message to the queue!
	err = channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}
}

func ReceiveDatasFromRabbitMQ(w http.ResponseWriter, r *http.Request) {

	//Subscribe to RabbitMQ
	msgs, err := channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	<-forever
}
