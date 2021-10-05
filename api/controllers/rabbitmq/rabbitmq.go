package rabbitmq_controller

import (
	"encoding/json"
	"fmt"
	"golang-rest-api/connections/rabbitmq"
	"net/http"

	"github.com/streadway/amqp"
)

func RabbitmqSendDataToQueue(w http.ResponseWriter, r *http.Request) {

	_, err := rabbitmq.RabbitChannel.QueueDeclare(
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
	err = rabbitmq.RabbitChannel.Publish(
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data Sended To Queue !",
	})
}

func ReceiveDatasFromRabbitMQ(w http.ResponseWriter, r *http.Request) {

	//Subscribe to RabbitMQ
	msgs, err := rabbitmq.RabbitChannel.Consume(
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data Got From Queue !",
	})

	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	<-forever

}
