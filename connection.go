package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

var (
	rabbitURL = "amqp://guest:guest@localhost:5672/"
)

// tao connection
func connectRabbitMQ() (*amqp.Connection, error) {
	return amqp.Dial(rabbitURL)
}

// reconnection lai neu mat ket noi
func reconnectRabbitMQ() *amqp.Connection {
	for {
		conn, err := connectRabbitMQ()
		if err == nil {
			return conn
		}
		fmt.Println("failed to connect. retrying in 2 seconds")
		time.Sleep(2 * time.Second)
	}
}
