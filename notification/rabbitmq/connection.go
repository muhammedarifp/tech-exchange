package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqConnection struct {
	Connection *amqp.Connection
}

func NewRabbitmqConnection() (*RabbitmqConnection, error) {
	conn, connErr := amqp.Dial("rabbitmq-service:5672")
	if connErr != nil {
		log.Fatalf("rabbitmq connection error found : %v", connErr)
		return &RabbitmqConnection{Connection: nil}, connErr
	}

	return &RabbitmqConnection{
		Connection: conn,
	}, nil
}
