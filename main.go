package main

import (
	"fmt"
	"time"
)

func main() {
	r := &RabbitMQ{}
	err := r.connectRabbitMQ()
	if err != nil {
		panic(err)
	}
	defer r.Close()

	done := make(chan bool)
	go produce(r, done)
	go consume(r, done)

	go func() {
		for {
			select {
			case <-done:
				return
			case err := <-r.check:
				fmt.Println("error received, reconnecting...", err)
				r.reconnectRabbitMQ()
			}
		}
	}()

	time.Sleep(10 * time.Second)
	close(done)
	fmt.Println("done!")
}
