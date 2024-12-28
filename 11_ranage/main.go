// 1. Range over an Array or Slice
// When you use range over an array or slice, it returns two values:

// The index (or position) of the element
// A copy of the value at that index (not the original value in the array or slice)

package main

import "fmt"

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}

// 	for i, value := range arr {
// 		fmt.Printf("Index: %d, Value: %d\n", i, value) // this syntax related to c language
// 	}
// }

// 2. Range over a Map
// When using range over a map, it returns two values:

// The key
// The value associated with that key

// func main() {
// 	myMap := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}

// 	for key, value := range myMap {
// 		fmt.Printf("Key: %s, Value: %d\n", key, value)
// 	}
// }

// 3. Range over a Channel
// When using range over a channel, it returns values sent over the channel until it is closed.

// func main() {
// 	ch := make(chan int, 3)
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3
// 	close(ch)

// 	for value := range ch {
// 		fmt.Println(value)
// 	}
// }

// 4. Range with only the Index
// If you only need the index and not the value,
//  you can omit the second variable by using the blank identifier _.

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}

// 	for i := range arr {
// 		fmt.Printf("Index: %d\n", i)
// 	}
// }

// 5. Range with only the Value
// If you only need the value and not the index,
//  you can omit the first variable by using the blank identifier _.

// func main() {
// 	slice := []string{"apple", "banana", "cherry"}

// 	for _, value := range slice {
// 		fmt.Printf("Value: %s\n", value)
// 	}
// }

// 6. Range over a String (Unicode Code Points)
// When using range over a string,
//  it returns the index and the Unicode code point (character) at that index.

// func main() {
// 	str := "hello"

// 	for i, ch := range str {
// 		fmt.Printf("Index: %d, Character: %c\n", i, ch)
// 	}
// }

// 7. Range over a Slice of Structs
// You can also use range over slices of structs. It works similarly to other data structures,
//  providing the index and a copy of the struct.

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, Name: %s, Age: %d\n", i, person.Name, person.Age)
	}
}
