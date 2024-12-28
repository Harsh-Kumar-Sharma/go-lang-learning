package main

import "fmt"

func main() {
	// age := 10

	// if age >= 18 {
	// 	fmt.Println("Person is adult")
	// } else {
	// 	fmt.Println("Person is not adult")
	// }

	// if age >= 18 {
	// 	fmt.Println("Person is adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is a teenager")
	// } else {
	// 	fmt.Println("Person is a Kid")
	// }

	// we are declare variable inside if construct

	if age := 20; age >= 18 {
		fmt.Println("Person is adult")
	} else if age >= 12 {
		fmt.Println("Person is a teenager")
	} else {
		fmt.Println("Person is a Kid")
	}

	// go does not have ternary operator ,you will use normal if else condation
}
