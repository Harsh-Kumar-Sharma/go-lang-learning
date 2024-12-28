package main

import "fmt"

// 1. Basic Map with String Keys and String Values
// This is the most common map type, where both the keys and values are strings.

// func main() {
// 	// Creating a map with string keys and string values
// 	studentGrades := map[string]string{
// 		"Alice":   "A",
// 		"Bob":     "B",
// 		"Charlie": "C",
// 	}

// 	// Accessing a value using a key
// 	fmt.Println("Alice's grade:", studentGrades["Alice"])
// }

// 2. Map with Integer Keys and String Values
// In this example, the keys are integers, and the values are strings.

// func main() {
// 	// Creating a map with integer keys and string values
// 	roomAssignments := map[int]string{
// 		101: "John",
// 		102: "Jane",
// 		103: "Doe",
// 	}

// 	// Accessing a value using a key
// 	fmt.Println("Room 102 is assigned to:", roomAssignments[102])
// }

// 3. Map with String Keys and Integer Values
// In this case, the keys are strings, and the values are integers.

// func main() {
// 	// Creating a map with string keys and integer values
// 	productPrices := map[string]int{
// 		"Apple":  2,
// 		"Banana": 1,
// 		"Orange": 3,
// 	}

// 	// Accessing a value using a key
// 	fmt.Println("Price of Banana:", productPrices["Banana"])
// }

// 4. Map with Struct Values
// Maps can also hold more complex values, such as structs. Here’s an example where the values are structs.

// Defining a struct for a book
// type Book struct {
// 	Title  string
// 	Author string
// 	Price  float64
// }

// func main() {
// 	// Creating a map with string keys and Book struct values
// 	books := map[string]Book{
// 		"Go Programming": {"Go Programming", "John Doe", 29.99},
// 		"Learn Go":       {"Learn Go", "Jane Smith", 19.99},
// 	}

// 	// Accessing a value (Book) using a key
// 	fmt.Println("Details of 'Go Programming':", books["Go Programming"])
// }

// 5. Map with Slice Values
// A map can store slices as values. This is useful when you need to associate a key with a list of values.

// func main() {
// 	// Creating a map with string keys and slice values
// 	courseStudents := map[string][]string{
// 		"Math":    {"Alice", "Bob", "Charlie"},
// 		"Science": {"Dave", "Eve", "Frank"},
// 	}

// 	science := courseStudents["Science"]

// 	// Accessing a slice using a key
// 	fmt.Println("Students in Science course:", science[1:])
// }

// 6. Map with Boolean Values
// Maps can have boolean values. This is useful when tracking true/false states.

// func main() {
// 	// Creating a map with string keys and boolean values
// 	taskCompletion := map[string]bool{
// 		"Task1": true,
// 		"Task2": false,
// 		"Task3": true,
// 	}

// 	// Checking if a task is completed
// 	fmt.Println("Is Task2 completed?", taskCompletion["Task2"])
// }

// 7. Empty Map
// An empty map is a map that has been declared but not initialized yet.
//  It’s useful for later population or conditional map creation.

// func main() {
// 	// Declaring an empty map with string keys and integer values
// 	var cityPopulations map[string]int

// 	// Initializing the map if it's nil
// 	if cityPopulations == nil {
// 		cityPopulations = make(map[string]int)

// 	}

// 	// Adding data to the map
// 	cityPopulations["New York"] = 8419600
// 	cityPopulations["Los Angeles"] = 3980400

// 	fmt.Println("City Populations:", cityPopulations)
// }

// 8. Map with Multiple Data Types in a Struct
// Go allows you to mix different data types in a struct, which can then be used as values in a map.

// Defining a struct with mixed data types
type Product struct {
	Name    string
	Price   float64
	InStock bool
}

func main() {
	// Creating a map with string keys and Product struct values
	// products := map[string]Product{
	// 	"Laptop": {Name: "Laptop", Price: 1200.00, InStock: true},
	// 	"Phone":  {Name: "Phone", Price: 800.00, InStock: false},
	// }

	var products = make(map[string]Product)

	products["Laptop"] = Product{Name: "Laptop", Price: 1200.00, InStock: true}

	// Accessing a value (Product) using a key
	fmt.Println("Product details of Laptop:", products)

}
