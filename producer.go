package main

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// chay 1 goroutine gui message toi rabbitMQ
func produce(r *RabbitMQ, done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			err := r.createQueue()
			if err != nil {
				fmt.Println("queue is not open", err)
			}
			err = r.ch.Publish( //gui tin nhan
				"", //exchange
				"testqueue",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte("hello world"),
				},
			)
			if err != nil {
				fmt.Println("fail", err)
			} else {
				fmt.Println("successfully published message to queue")
			}
			time.Sleep(1 * time.Second)
		}
	}
}
