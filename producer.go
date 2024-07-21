package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

// chay 1 goroutine gui message toi rabbitMQ
func produce(wg *sync.WaitGroup, done <-chan bool) {
	defer wg.Done()

	for {
		select {
		case <-done: //dung khi done
			return
		default:
			conn := reconnectRabbitMQ()
			ch, err := conn.Channel()
			if err != nil {
				fmt.Println(err)
			}
			q, err := ch.QueueDeclare( //tao hang doi testqueue neu ch ton tai
				"testqueue",
				false,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(q)

			err = ch.Publish( //gui tin nhan
				"",
				"testqueue",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte("hello world"),
				},
			)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("successfully published message to queue")
			}

			ch.Close()
			conn.Close()
			time.Sleep(1 * time.Second)
		}

	}
}
