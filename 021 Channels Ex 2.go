// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffered channel so no send ever blocks. Don't allocate more capacity
// than you need. Have the main goroutine store each random number it receives
// in a slice. Print the slice values then terminate the program.
package main

// Add imports.
import (
	"fmt"
	"math/rand"
)

// Declare constant for number of goroutines.
const amount = 100

func main() {
	// Create the buffered channel with room for
	// each goroutine to be created.
	ch := make(chan int, amount)

	// Iterate and launch each goroutine.
	for range amount {
		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			ch <- rand.Intn(amount)
		}()
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	wait := amount

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	var nums []int
	for wait > 0 {
		d := <-ch
		nums = append(nums, d)
		wait--
	}

	// Print the values in our slice.
	fmt.Println("Result: ", nums)
}
