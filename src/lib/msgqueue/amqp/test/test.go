package main

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file:", err)
	// }

	amqpURL := os.Getenv("AMQP_URL")

	//fmt.Println("url: ", amqpURL)

	// if amqpURL == "" {
	// 	amqpURL = "amqp://guest:guest@localhost:5672"
	// }

	connection, err := amqp.Dial(amqpURL)
	if err != nil {
		fmt.Println("Error connecting to RabbitMQ:", err.Error())
	}

	channel, err := connection.Channel()
	if err != nil {
		fmt.Println("Could not open channel:", err.Error())
	}

	err = channel.ExchangeDeclare("event", "topic", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	message := amqp.Publishing{
		Body: []byte("Hello World"),
	}
	err = channel.Publish("event", "some-routing-key", false, false, message)
	if err != nil {
		panic("error while publishing message:" + err.Error())
	}

	//Queue
	_, err = channel.QueueDeclare("my_queue", false, false, false, false, nil)
	if err != nil {
		panic("error while declaring queue:" + err.Error())
	}

	err = channel.QueueBind("my_queue", "#", "event", false, nil)
	if err != nil {
		panic("error while binding the queue:" + err.Error())
	}

	msgs, err := channel.Consume("my_queue", "", true, false, false, false, nil)
	if err != nil {
		panic("error while consuming the queue:" + err.Error())
	}

	for msg := range msgs {
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(false)
	}

}
