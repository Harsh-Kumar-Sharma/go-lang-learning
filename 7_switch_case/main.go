// 1 Basic Switch: Matches a single value to multiple cases.
// 2 Switch Without Expression: Implies true for each case, allowing complex conditions.
// 3 Multiple Expressions in a Case: Allows a single case to match multiple values.
// 4 Type Switch: Matches the type of an interface value.
// 5 Fallthrough: Allows control over whether to continue to the next case.
// 6 Switch with Conditions: Allows condition-based matching.

// 1. Basic Switch Case
// A basic switch statement evaluates an expression and matches it against the
// case values. The first matching case is executed.

package main

import "fmt"

// func main() {
// 	day := 17
// 	switch day {
// 	case 1:
// 		fmt.Println("Monday")
// 	case 2:
// 		fmt.Println("Tuesday")
// 	case 3:
// 		fmt.Println("Wednesday")
// 	case 4:
// 		fmt.Println("Thursday")
// 	case 5:
// 		fmt.Println("Friday")
// 	default:
// 		fmt.Println("Weekend")
// 	}
// }

// 2. Switch Without Expression (Implicit true)
// Go allows you to write a switch without an expression, which behaves like an implicit switch true.
//  This is useful for testing multiple conditions.

// func main() {
// 	x := 7
// 	switch {
// 	case x < 5:
// 		fmt.Println("x is less than 5")
// 	case x > 5 && x < 10:
// 		fmt.Println("x is between 5 and 10")
// 	default:
// 		fmt.Println("x is greater than or equal to 10")
// 	}
// }

// 3. Switch with Multiple Expressions (Multiple Cases)
// You can have a case with multiple values, where if any of the values match the expression,
//  that case will be executed.

// func main() {
// 	fruit := "apple"
// 	switch fruit {
// 	case "apple", "banana", "cherry":
// 		fmt.Println("It's a fruit.")
// 	case "carrot", "potato":
// 		fmt.Println("It's a vegetable.")
// 	default:
// 		fmt.Println("Unknown.")
// 	}
// }

// 4. Switch with Type Assertion (Type Switch)
// Go allows type switches to work with interface types,
//  which is useful when you need to handle different types that implement an interface.

// func checkType(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println("Integer:", v)
// 	case string:
// 		fmt.Println("String:", v)
// 	case bool:
// 		fmt.Println("Boolean:", v)
// 	default:
// 		fmt.Println("Unknown type")
// 	}
// }

// func checkType(i interface{}) {
// 	switch i.(type) {
// 	case int:
// 		fmt.Println("Integer")
// 	case string:
// 		fmt.Println("String")
// 	case bool:
// 		fmt.Println("Boolean")
// 	default:
// 		fmt.Println("Unknown type")
// 	}
// }

// func main() {
// 	checkType(42)
// 	checkType("hello")
// 	checkType(true)
// 	checkType(3.14)
// }

// 5. Fallthrough in Switch
// By default, Go does not allow the execution of subsequent
//  case blocks after a match (like break in C or Java). However,
//   you can use fallthrough to explicitly continue to the next case.

// func main() {
// 	number := 1
// 	switch number {
// 	case 1:
// 		fmt.Println("One")
// 		fallthrough
// 	case 2:
// 		fmt.Println("Two")
// 		fallthrough
// 	case 3:
// 		fmt.Println("Three")
// 	default:
// 		fmt.Println("Unknown number")
// 	}
// }

// 6. Switch with Conditions (Comparing Expressions)
// You can also use conditions in case statements to make more complex comparisons.

func main() {
	number := 5
	switch {
	case number%2 == 0:
		fmt.Println("Even number")
	case number%2 != 0:
		fmt.Println("Odd number")
	}
}
