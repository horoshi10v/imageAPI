package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
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

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			createTempFile(d.Body, d.MessageId)
		}
	}()
	log.Println("Successfully Connected to our RabbitMQ Instance")
	log.Println(" [*] - Waiting for messages")
	<-forever
}

func createTempFile(fileBytes []byte, filename string) {
	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	//tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	tempFile, err := os.Create("upload-images/" + filename)
	if err != nil {
		log.Println(err)
	}
	defer tempFile.Close()
	// write this byte array to our temporary file
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		log.Fatal("can't write to temporary file", err)
	}
	// return that we have successfully uploaded our file!
	log.Printf("Successfully Uploaded File\n")
}
