// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

func main() {

	// Declare a nil slice of integers.
	var numbers []int

	// Append numbers to the slice.
	for i := range 10 {
		numbers = append(numbers, i*100)
	}

	// Display each value in the slice.
	for _, v := range numbers {
		fmt.Printf("<%d>\n", v)
	}

	// Declare a slice of strings and populate the slice with names.
	names := []string{"a", "b", "c", "d", "e"}

	// Display each index position and slice value.
	for k := range names {
		fmt.Printf("%d: %q\n", k, names[k])
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	names_1_2 := names[1:3]

	// Display each index position and slice values for the new slice.
	for n, name := range names_1_2 {
		fmt.Printf("%d: %q\n", n, name)
	}
}
