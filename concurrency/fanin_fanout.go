package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker simulates a task performed by a worker goroutine
func worker(id int, input <-chan int, output chan<- int) {
	for num := range input {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) // Simulate processing time
		fmt.Printf("Worker %d processed: %d\n", id, num)
		output <- num * num // Processed result (square of the number)
	}
}

// FanOut distributes tasks to multiple workers
func fanOut(inputs []int, workers int) <-chan int {
	inputChannel := make(chan int, len(inputs))  // Input channel
	outputChannel := make(chan int, len(inputs)) // Output channel

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		go worker(i+1, inputChannel, outputChannel)
	}

	// Send inputs to the inputChannel
	go func() {
		for _, input := range inputs {
			inputChannel <- input
		}
		close(inputChannel) // Close the input channel to signal workers we're done
	}()

	return outputChannel
}

// FanIn collects results from the output channel
func fanIn(outputChannel <-chan int, size int) []int {
	results := make([]int, 0, size)
	for i := 0; i < size; i++ {
		results = append(results, <-outputChannel)
	}
	return results
}

func FaninFanout() {
	rand.Seed(time.Now().UnixNano()) // Seed for random numbers

	// Inputs
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Fan-Out: Distribute work to 3 workers
	outputChannel := fanOut(numbers, 3)

	// Fan-In: Collect results
	results := fanIn(outputChannel, len(numbers))

	// Print final results
	fmt.Println("Final Results:", results)
}
