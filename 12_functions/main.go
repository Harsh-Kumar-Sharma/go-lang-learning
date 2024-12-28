// A function in Go is defined with the following syntax:

// func FunctionName(parameter1 Type, parameter2 Type) ReturnType {
//     // Function body
//     return value
// }

package main

import "fmt"

// 1. Basic Function Definition
// A basic function in Go is defined using the func keyword, followed by the function name, parameters, and return type.

// Example:

// Function definition
// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	result := add(3, 4)
// 	fmt.Println(result) // Output: 7
// }

// 2. Function with Multiple Return Values
// Go functions can return multiple values, which is a useful feature for cases like returning both a result and an error.

// Example:

// Function with multiple return values
// func divide(a, b int) (int, int) {
// 	return a / b, a % b
// }

// func main() {
// 	quotient, remainder := divide(10, 3)
// 	fmt.Println("Quotient:", quotient)   // Output: Quotient: 3
// 	fmt.Println("Remainder:", remainder) // Output: Remainder: 1
// }

// 3. Anonymous Functions (Lambdas)
// An anonymous function is a function without a name, often used for short, one-time use cases.

// Example:

// func main() {
// 	// Anonymous function
// 	result := func(x, y int) int {
// 		return x + y
// 	}(3, 4)

// 	fmt.Println(result) // Output: 7
// }

// 4. Higher-Order Functions
// A higher-order function is a function that takes other functions as arguments or returns them as results.

// Example:

// Higher-order function
// func applyFunc(f func(int, int) int, a int, b int) int {
// 	return f(a, b)
// }

// func add(x, y int) int {
// 	return x + y
// }

// func sub(x, y int) int {
// 	return x - y
// }

// func divide(x, y int) int {
// 	return x / y
// }

// func multiple(x, y int) int {
// 	return x * y
// }

// func main() {
// 	result := applyFunc(add, 5, 3)
// 	result1 := applyFunc(sub, 5, 3)
// 	result2 := applyFunc(divide, 5, 3)
// 	result3 := applyFunc(multiple, 5, 3)
// 	fmt.Println(result, result1, result2, result3) // Output: 8
// }

// 5. Variadic Functions
// A variadic function is a function that can accept a variable number of arguments.

// Example:

// Variadic function
// func sum(numbers ...int) int {
// 	total := 0
// 	for _, num := range numbers {
// 		total += num
// 	}
// 	return total
// }

// func main() {
// 	arr := []int{2, 2, 2, 2, 2}
// 	fmt.Println(sum(1, 2, 3))       // Output: 6
// 	fmt.Println(sum(1, 2, 3, 4, 5)) // Output: 15
// 	fmt.Println(sum(arr...))        // Output: 10
// }

// 6. Named Return Values
// Go functions can have named return values,
// which allows you to specify the names of the values you are returning.
// These names can be used in the body of the function and automatically returned at the end.

// Example:

// Function with named return values
// func divide(a, b int) (quotient, remainder, sum int) {
// 	quotient = a / b
// 	remainder = a % b
// 	sum = a + b
// 	return // Named return values are automatically returned
// }

// func main() {
// 	q, r, s := divide(10, 3)
// 	fmt.Println("Quotient:", q)  // Output: Quotient: 3
// 	fmt.Println("Remainder:", r) // Output: Remainder: 1
// 	fmt.Println("Sum:", s)       // Output: Sum: 13
// }

// 7. Defer, Panic, and Recover
// Go has built-in support for handling errors and deferring actions using defer, panic, and recover.

// defer: Executes a function call just before the surrounding function returns.
// panic: Stops execution and starts unwinding the stack.
// recover: Handles a panic and prevents the program from crashing.
// Example:

// func causePanic() {
// 	panic("Something went wrong!")
// }

// func safeFunction() {
// 	defer fmt.Println("This will be printed after panic.")
// 	causePanic()
// }

// func main() {
// 	// Recover from panic
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from:", r)
// 		}
// 	}()

// 	safeFunction() // Output: Recovered from: Something went wrong!
// }

// 8. Methods
// In Go, methods are functions that are associated with a type (usually a struct). Methods allow you to define behavior on that type.

// Example:

type Rectangle struct {
	width, height int
}

// Method for the Rectangle type
func (r Rectangle) Area() int {
	return r.width * r.height
}

func main() {
	rect := Rectangle{width: 5, height: 6}
	fmt.Println("Area of rectangle:", rect.Area()) // Output: Area of rectangle: 30
}
