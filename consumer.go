package main

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// chay 1 goroutine, nhan message tu rabbitmq
func consume(r *RabbitMQ, done chan bool) {
	for {
		err := r.createQueue()
		if err != nil {
			fmt.Println("create queue failed", err)
			r.check <- err
			time.Sleep(2 * time.Second)
			continue
		}
		msgs, err := r.ch.Consume(
			"testqueue",
			"",
			true,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			fmt.Println("consume fail", err)
			r.check <- err
			time.Sleep(2 * time.Second)
			continue
		}
		go func() { //su dung goroutine de xu ly tin nhan
			for {
				select {
				case <-done:
					return
				case d := <-msgs:
					fmt.Printf("Received a message: %s\n", d.Body)
				}
			}
		}()
		select {
		case <-done:
			r.Close()
			return
		case err := <-r.conn.NotifyClose(make(chan *amqp.Error)):
			fmt.Println(err)
			r.Close()
			r.check <- err
			continue
		}
	}

}
