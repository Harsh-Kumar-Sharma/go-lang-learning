// In Go, a struct is a composite data type that groups together variables (fields) of different types
//  into a single unit. This allows you to model real-world entities and objects.
//   Structs are similar to classes in object-oriented languages, but they do not support
//    inheritance or methods directly on the struct type unless explicitly defined.

//Example:

package main

import "fmt"

// Define a struct named "Person"
// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// 	Email     string
// }

// func main() {
// 	// Initialize a struct with values
// Creating an Instance of a Struct:
// 	p1 := Person{
// 		FirstName: "John",
// 		LastName:  "Doe",
// 		Age:       30,
// 		Email:     "john.doe@example.com",
// 	}

// Accessing Struct Fields:
// You can access the fields of the struct using the dot (.) operator
// 	fmt.Println("First Name:", p1.FirstName)
// 	fmt.Println("Last Name:", p1.LastName)
// 	fmt.Println("Age:", p1.Age)
// 	fmt.Println("Email:", p1.Email)
// }

// Zero Value of a Struct:
// If you create a struct without initializing its fields, the fields will have the zero value for their respective types (e.g., "" for strings, 0 for integers, and nil for pointers).

// go
// Copy code
// var p2 Person  // Zero value struct
// fmt.Println(p2)  // Prints {"" "" 0 ""}

// 2.Working with Methods on Structs

// Define a struct named "Person"
// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// 	Email     string
// }

// // Method to print a greeting message
// func (p Person) Greet() {
// 	fmt.Println("Hello, my name is", p.FirstName, p.LastName)
// }

// // Method to return the age of a person
// func (p Person) GetAge() int {
// 	return p.Age
// }

// func main() {
// 	p1 := Person{
// 		FirstName: "John",
// 		LastName:  "Doe",
// 		Age:       30,
// 		Email:     "john.doe@example.com",
// 	}

// 	// Call the methods
// 	p1.Greet()                       // Outputs: Hello, my name is John Doe
// 	fmt.Println("Age:", p1.GetAge()) // Outputs: Age: 30
// }

// Structs vs. Pointers to Structs
// You can also pass pointers to structs to avoid copying the struct data
//  when working with large structs or
//  when you need to modify the struct inside a method

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *Person) CelebrateBirthday() {
	p.Age++
}

func main() {
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	p2 := Person{
		FirstName: "Harsh",
		LastName:  "Sharma",
		Age:       20,
	}

	fmt.Println("Before Birthday:", p1.Age) // Outputs: 30
	p1.CelebrateBirthday()
	fmt.Println("After Birthday:", p1.Age) // Outputs: 31

	fmt.Println("Before Birthday:", p2.Age) // Outputs: 20
	p2.CelebrateBirthday()
	fmt.Println("After Birthday:", p2.Age) // Outputs: 21
}
