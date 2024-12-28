// In Go, an interface is a type that specifies a set of methods but does not provide
//  the implementation for them. A type (such as a struct) implements an interface by providing
//   definitions for all the methods declared by the interface.

// Key Points about Interfaces:
// 1.Implicit Implementation: Unlike other languages like Java, Go does not require explicit
//  declarations of interface implementation. A type implicitly satisfies an interface if
//  it implements all the methods declared by the interface.
// 2.Empty Interface: An empty interface (interface{}) can hold any value since every type implements
// at least zero methods.

// Example of Interface in Go:

package main

import "fmt"

// Define the "Speaker" interface with a "Speak" method
// type Speaker interface {
// 	Speak() string
// }

// // Define a struct called "Person"
// type Person struct {
// 	Name string
// }

// // Define a method for the "Person" struct to satisfy the "Speaker" interface
// func (p Person) Speak() string {
// 	return "Hello, my name is " + p.Name
// }

// // Define a struct called "Robot"
// type Robot struct {
// 	ID string
// }

// // Define a method for the "Robot" struct to satisfy the "Speaker" interface
// func (r Robot) Speak() string {
// 	return "Beep boop, I am robot number " + r.ID
// }

// // Function that accepts any type that satisfies the "Speaker" interface
// func introduce(speaker Speaker) {
// 	fmt.Println(speaker.Speak())
// }

// func main() {
// 	// Create an instance of "Person" and "Robot"
// 	person := Person{Name: "Alice"}
// 	robot := Robot{ID: "R2-D2"}

// 	// Passing the "person" and "robot" objects to the "introduce" function
// 	introduce(person) // Output: Hello, my name is Alice
// 	introduce(robot)  // Output: Beep boop, I am robot number R2-D2
// }

// Empty Interface:
// Go also provides an empty interface, interface{}, which can hold values of any type.

func printAny(value interface{}) {
	fmt.Println(value)
}

func main() {
	printAny(42)          // integer
	printAny("Hello Go!") // string
	printAny(3.14)        // float
}
