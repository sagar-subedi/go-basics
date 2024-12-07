package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	bufferSize = 5  // Size of the buffer
	numItems   = 10 // Number of items to produce
)

// Shared buffer and synchronization primitives
var buffer = make(chan int, bufferSize) // A channel serves as the buffer
var wg sync.WaitGroup                   // WaitGroup to ensure all goroutines complete

func producer(id int) {
	defer wg.Done() // Decrement the WaitGroup counter when done

	for i := 0; i < numItems; i++ {
		item := rand.Intn(100) // Producing a random item
		buffer <- item         // Send item to the buffer (blocks if buffer is full)
		fmt.Printf("Producer %d produced: %d\n", id, item)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
}

func consumer(id int) {
	defer wg.Done() // Decrement the WaitGroup counter when done

	for i := 0; i < numItems; i++ {
		item := <-buffer // Receive item from the buffer (blocks if buffer is empty)
		fmt.Printf("Consumer %d consumed: %d\n", id, item)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
}

func main() {
	// Add producers and consumers to the WaitGroup
	wg.Add(2) // One producer and one consumer for this example

	go producer(1) // Start producer goroutine
	go consumer(1) // Start consumer goroutine

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All done!")
}
