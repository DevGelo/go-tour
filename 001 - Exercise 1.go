// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"
	
// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var x string
	var y int
	var z bool

	// Display the value of those variables.
	fmt.Printf("x : %T <%v>\n", x, x)
	fmt.Printf("y : %T <%v>\n", y, y)
	fmt.Printf("z : %T <%v>\n", z, z)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	xx := "Good morning!"
	yy := 42
	zz := true

	// Display the value of those variables.
	fmt.Printf("xx : %T <%v>\n", xx, xx)
	fmt.Printf("yy : %T <%v>\n", yy, yy)
	fmt.Printf("zz : %T <%v>\n", zz, zz)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Printf("pi : %T <%v>\n", pi, pi)
}
