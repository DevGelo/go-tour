// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Implement a generic function named copyfy that is constrained to only
// making copies of slices of type string or int.
package main

import "fmt"

// Declare an interface named copyer that creates a constraint on
// string and int.
type copyer interface {
	string | int
}

// Implement a generic function named copyfy that accepts a slice of some
// type T but constrained on the copyer interface.
func copyfy [T copyer] (source []T) []T {
	destination := make([]T, len(source))
	copy(destination, source)

	return destination
}

func main() {

	// Construct a slice of string with three values.
	s := []string{"a", "b", "c"}

	// Call the copyfy function to make a copy of the slice.
	cs := copyfy(s)

	// Display the slice and the copy.
	fmt.Println("Source: ", s)
	fmt.Println("Copy: ", cs)

	// Construct a slice of int with three values.
	n := []int{10, 20, 30}

	// Call the copyfy function to make a copy of the slice.
	cn := copyfy(n)

	// Display the slice and the copy.
	fmt.Println("Source: ", n)
	fmt.Println("Copy: ", cn)
}
