package main

import (
	"fmt"
	"pubsub/broker"
	"sync"
	"time"
)

func main() {
	topic := broker.NewTopic("topic", 10)

	var wg sync.WaitGroup

	fmt.Println("start multiple producers")
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			producer := topic.AddProducer()
			for j := 1; j <= 5; j++ {
				message := fmt.Sprintf("Producer %d message %d", id, j)
				producer <- message
				time.Sleep(time.Millisecond * 500)
			}
		}(i)
	}

	fmt.Println("Start multiple consumers")
	consumer1 := topic.AddConsumer("1")
	consumer2 := topic.AddConsumer("2")
	go func(a int) {
		for {
			msg := <-consumer1
			fmt.Printf("Consumer %d received: %s\n", a, msg)
		}
	}(1)
	go func(a int) {
		for {
			msg := <-consumer2
			fmt.Printf("Consumer %d received: %s\n", a, msg)
		}
	}(2)
	// go func() {
	// 	defer wg.Done()
	// 	consumer := topic.AddConsumer()
	// 	for msg := range consumer {
	// 		fmt.Printf("Consumer received: %s\n", msg)
	// 	}
	// }()
	wg.Wait()
}
