// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name string
	email string
	age int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	simon:= user{
		name: "Simon",
		email: "simon@mail.com",
		age: 42,
	}

	// Display the field values.
	fmt.Printf("%+v\n", simon)

	// Declare a variable using an anonymous struct.
	bill := struct{
		name string
		email string
		age uint8
	}{
		name: "Bill",
		email: "bill@mail.com",
		age: 35,
	}

	// Display the field values.
	fmt.Println("Name", bill.name)
	fmt.Println("Email", bill.email)
	fmt.Println("Age", bill.age)
}
