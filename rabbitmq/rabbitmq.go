package rabbitmq

import (
	"fmt"
	"net/http"

	"github.com/streadway/amqp"
)

type Rabbit struct {
	rabCon     *amqp.Connection
	rabChannel *amqp.Channel
}

type RabbitMQFunctions struct {
}

func RabbitMQConnect() *Rabbit {

	amqpp, err := amqp.Dial("amqps://oryeecvq:Ba5zvf5J9JD5tObauu054nnssDfEJPXh@cow.rmq2.cloudamqp.com/oryeecvq")

	if err != nil {
		fmt.Printf("RabbitMQ Connection is not Running !\n")
	} else {
		fmt.Printf("RabbitMQ Connection is Running !\n")
	}

	//Create a channel
	ch, err := amqpp.Channel()

	return &Rabbit{
		rabCon:     &amqp.Connection{},
		rabChannel: ch,
	}

	//defer ch.Close()
}

func (rab *Rabbit) RabbitmqSendDataToQueue(http.ResponseWriter, *http.Request) {

	_, err := rab.rabChannel.QueueDeclare(
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
	err = rab.rabChannel.Publish(
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

func (rab *Rabbit) ReceiveDatasFromRabbitMQ(http.ResponseWriter, *http.Request) {

	//Subscribe to RabbitMQ
	msgs, err := rab.rabChannel.Consume(
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
