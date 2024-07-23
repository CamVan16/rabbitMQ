package main

import (
	amqp "github.com/rabbitmq/amqp091-go" //rabbitmq
)

var (
	rabbitURL = "amqp://guest:guest@localhost:5672/"
)

type RabbitMQ struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
	check bool
}

// tao connection, channel
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
	r.setCheck(true)
	return nil
}

func (r *RabbitMQ) getCheck() bool {
	return r.check
}

func (r *RabbitMQ) setCheck(c bool) {
	r.check = c
}

// tao queue
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

// dong ket noi
func (r *RabbitMQ) Close() {
	r.check = false
	r.ch.Close()
	r.conn.Close()

}
