// A closure in Go is a function value that "captures" variables from its surrounding environment.
//When you define a function inside another function, the inner function can access and modify
// variables from the outer function, even after the outer function has finished executing.
// This is called a closure.

// Key Concepts:
// 1.Inner Function: The function defined inside another function.
// 2.Captured Variables: The outer function's variables that are accessible inside the inner function.
// 3.Persistence: The inner function retains access to the captured
//variables even after the outer function has returned.
// 5.Example of Closure in Go:

package main

import "fmt"

// Outer function that returns an inner function
func outer() func() int {
	// Local variable in the outer function
	counter := 0

	// Inner function (closure)
	return func() int {
		counter++
		return counter
	}
}

func main() {
	// Create a closure by calling outer()
	increment := outer()

	// Call the closure multiple times
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3

	// Create another closure
	anotherIncrement := outer()
	fmt.Println(anotherIncrement()) // 1
}

// Explanation:
//1. outer() function: This function defines a local variable counter and returns an anonymous
// function (a closure). The anonymous function has access to the counter variable
// from the outer() function.

// 2.inner function: The returned function return func() int is the closure.
//  It increments the counter variable each time it is called,
//  demonstrating how the closure "captures" the variable from the outer function and continues
//  to update it across function calls.

// 3.Capturing state: When you call increment(), the state of the counter variable is preserved.
//  Each time increment() is invoked, it continues from the last value of counter,
//  even though the outer() function has finished executing.

// 4.Multiple closures: A new closure, anotherIncrement, is created when calling outer() again.
//  This creates a new instance of the counter variable,
//  which starts at 0, hence anotherIncrement() starts from 1.

// Key Points:
// 1.Capturing: Closures capture variables from their surrounding environment.
// 2.State retention: The closure "remembers" the state of variables from the outer function.
// 3.Multiple instances: Each closure can have its own independent state
//  if the outer function is called multiple times.
// Closures are commonly used in Go for scenarios like callbacks, event handling,
//  and maintaining state in a function without needing to create complex data structures.
