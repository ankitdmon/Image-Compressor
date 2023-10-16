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

func PublishToRabbitMQ(message string) error {
	connection, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         []byte(message),
	})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	log.Printf("Published a message: %s", message)

	return nil
}
