// pointers are variables that store the memory address of another variable.
// Instead of holding a direct value, a pointer holds the address of a variable.
// Using pointers allows you to modify the original value of a variable from within a
// function and can be more efficient for large data structures because it avoids copying the data.

package main

import "fmt"

// func main() {
// 	// Declare a variable
// 	var a int = 58

// 	// Declare a pointer that holds the address of variable 'a'
// 	var ptr *int = &a

// 	fmt.Println("Value of a:", a)           // prints 58
// 	fmt.Println("Address of a:", &a)        // prints the memory address of a
// 	fmt.Println("Value of ptr:", ptr)       // prints the memory address of 'a'
// 	fmt.Println("Dereferencing ptr:", *ptr) // prints the value at the address stored in ptr, which is 58

// 	// Change the value of a using pointer
// 	*ptr = 100
// 	fmt.Println("New value of a after dereferencing pointer:", a) // prints 100
// }

//Example with Functions

func increment(x *int) {
	*x++ // Dereference the pointer and increment the value at that memory address
}

func main() {
	num := 10
	fmt.Println("Before increment:", num) // 10
	increment(&num)
	fmt.Println("After increment:", num) // 11
}
