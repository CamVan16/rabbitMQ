package main

import (
	"fmt"
	"time"
)

func main() {
	r := &RabbitMQ{}
	err := r.connectRabbitMQ()
	if err != nil {
		fmt.Println(err)
	}
	defer r.Close()

	done := make(chan bool)
	go produce(r, done)
	go consume(r, done)
	//sd go de reconnect
	go func() {
		c := make(chan bool)
		c <- r.getCheck()
		for {
			select {
			case <-done:
				return
			case <-c:
				fmt.Println("reconnecting...", err)
				r.connectRabbitMQ()
			}
		}
	}()

	time.Sleep(10 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("done!")

}
