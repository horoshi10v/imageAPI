package rabbit

import (
	helper "api/helpers"
	"fmt"
	"github.com/streadway/amqp"
)

// ConsumeMsgRabbit consume messages from RabbitMQ
func ConsumeMsgRabbit() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	for d := range msgs {
		helper.Resize(d.Body, d.MessageId)
	}

}
