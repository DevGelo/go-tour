// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

// Add imports.
import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {

	// Create the channel for sharing results.
	ch := make(chan int)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	shutdown := make(chan struct{})

	// Define the size of the worker pool. Use runtime.GOMAXPROCS(0) to size the pool based on number of processors.
	poolSize := runtime.GOMAXPROCS(0)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(poolSize)

	// Create a fixed size pool of goroutines to generate random numbers.
	for i := range poolSize {
		go func(id int) {
			fmt.Printf("[%d] worker began\n", id)
			defer fmt.Printf("[%d] worker finished\n", id)
			
			// Start an infinite loop.
			for
			{
				// Generate a random number up to 1000.
				d := rand.Intn(1000)

				// Use a select to either send the number or receive the shutdown signal.
				select
				{
					// In one case send the random number.
					case ch <- d:
						fmt.Printf("[%d] worker sent %d\n", id, d)

					// In another case receive from the shutdown channel.
					case <- shutdown:
						wg.Done()
						fmt.Printf("[%d] worker shutdown\n", id)
					return
				}
			}
		}(i)
	}

	// Create a slice to hold the random numbers.
	var nums []int

	// Receive from the values channel with range.
	for d := range ch {
		fmt.Println("received number ", d)

		// continue the loop if the value was even.
		if d % 2 == 0 {
			fmt.Println("Discarded ", d)
			continue
		}

		// Store the odd number.
		nums = append(nums, d)

		// break the loop once we have 100 results.
		if len(nums) == 100 {
			break
		}
	}

	// Send the shutdown signal by closing the shutdown channel.
	close(shutdown)
	fmt.Println("send shutdown signal")

	// Wait for the Goroutines to finish.
	fmt.Println()
	fmt.Println("waiting workers")
	wg.Wait()

	// Print the values in our slice.
	fmt.Println()
	fmt.Println("Result count: ", len(nums))
	fmt.Println("Result: ", nums)
}
