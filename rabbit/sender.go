package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// SendMsgRabbit Send []byte to RabbitMQ consumer
func SendMsgRabbit(fileBytes []byte, filename string) {
	//TODO move connection to .env
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Failed Initializing Broker Connection", err)
	}
	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established
	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	log.Println(q)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		log.Println(err)
	}

	// attempt to publish a message to the queue
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        fileBytes,
			MessageId:   filename,
		},
	)

	if err != nil {
		fmt.Println(err)
	}
	log.Println("Successfully Published Message to Queue")

}
