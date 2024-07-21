package main

import (
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

// chay 1 goroutine, nhan message tu rabbitmq
func consume(wg *sync.WaitGroup, done <-chan bool) {
	defer wg.Done()

	for {
		select {
		case <-done:
			return
		default:
			conn := reconnectRabbitMQ()
			ch, err := conn.Channel()
			if err != nil {
				fmt.Println(err)
				conn.Close()
				continue
			}

			_, err = ch.QueueDeclare( //tao hang doi trc khi tieu thu message
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
			msgs, err := ch.Consume(
				"testqueue",
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				fmt.Println(err)
				ch.Close()
				conn.Close()
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
				ch.Close()
				conn.Close()
				return
			case err := <-conn.NotifyClose(make(chan *amqp.Error)):
				fmt.Println(err)
				ch.Close()
				conn.Close()
				continue
			}
		}
	}

}
