package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg+" : %v", err)
	}
}

func NewRabbitmqConnection() (*amqp.Connection, error) {
	conn, connErr := amqp.Dial("amqp://localhost:5672")
	failOnError(connErr, "connection error")

	return conn, nil
}
