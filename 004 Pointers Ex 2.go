// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct{
	name string
	count int
}

// Create a function that changes the value of one of the user fields.
func funcName(user *user, value int/* add pointer parameter, add value parameter */ ) {

	// Use the pointer to change the value that the
	// pointer points to.
	user.count = value
}

func main() {

	// Create a variable of type user and initialize each field.
	bill := user{
		name: "bill",
		count: 1,
	}

	// Display the value of the variable.
	fmt.Printf("%+v\n", bill)

	// Share the variable with the function you declared above.
	funcName(&bill, 12)

	// Display the value of the variable.
	fmt.Printf("%+v\n", bill)
}
