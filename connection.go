package main

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go" //rabbitmq
)

var (
	rabbitURL = "amqp://guest:guest@localhost:5672/"
)

type RabbitMQ struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
	check chan error
}

// tao connection
func (r *RabbitMQ) connectRabbitMQ() error {
	var err error
	r.conn, err = amqp.Dial(rabbitURL)
	if err != nil {
		return err
	}
	r.ch, err = r.conn.Channel()
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQ) createQueue() error {
	var err error
	r.queue, err = r.ch.QueueDeclare(
		"testqueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

// reconnection lai neu mat ket noi
func (r *RabbitMQ) reconnectRabbitMQ() error {
	for {
		err := r.connectRabbitMQ()
		if err == nil {
			return err
		}
		fmt.Println("failed to connect. retrying in 2 seconds")
		time.Sleep(2 * time.Second)
	}
}

// dong ket noi
func (r *RabbitMQ) Close() {
	r.ch.Close()
	r.conn.Close()

}
