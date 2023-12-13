package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateRabbitMqConnection() (*amqp.Connection, error) {
	conn, connErr := amqp.Dial("amqp://localhost:5672")
	if connErr != nil {
		return nil, connErr
	}

	return conn, nil
}
