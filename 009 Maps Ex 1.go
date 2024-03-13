// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import "fmt"

func main() {

	// Declare and make a map of integer type values.
	counts := map[string]int{}

	// Initialize some data into the map.
	counts["cars"] = 1
	counts["apples"] = 3
	counts["exames"] = 2
	counts["days"] = 7
	counts["months"] = 12

	// Display each key/value pair.
	for key, value := range counts {
		fmt.Printf("%v: %d\n", key, value)
	}
}
