package messaging

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const (
	rabbitMQURL = "amqp://guest:guest@localhost:5672/"
	queueName   = "processImageQueue"
)

func ConsumeFromRabbitMQ() {
	connection, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Println("Waiting for messages. To exit, press Ctrl+C")

	for msg := range msgs {
		message := string(msg.Body)
		fmt.Printf("Received a message: %s\n", message)
	}
}
