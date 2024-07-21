package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup //tao wg de doi cac goroutines hoan thanh
	done := make(chan bool)

	wg.Add(2)

	go produce(&wg, done)
	go consume(&wg, done)

	time.Sleep(5 * time.Second)

	close(done)

	wg.Wait() // doi ca 2 goroutine hoan thanh
	fmt.Println("done!")
}
