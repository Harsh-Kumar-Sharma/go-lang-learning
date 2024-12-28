// 1. Basic Array
// A basic array in Go is declared with a fixed size, and the size is part of its type.
// The elements in the array are stored contiguously in memory.

package main

import "fmt"

// func main() {
// 	var arr [5]int // Declaring an array of 5 integers
// 	arr[0] = 10
// 	arr[1] = 20
// 	arr[2] = 30
// 	arr[3] = 40
// 	arr[4] = 50

// 	fmt.Println("Array:", arr)
// }

// 2. Array with Initialization
// You can initialize an array at the time of declaration by specifying its values directly.

// func main() {
// 	var test = [3]int{1, 2, 3}

// 	fmt.Println(test)

// 	arr := [5]int{1, 2, 3, 4, 5} // Array initialization
// 	fmt.Println("Array:", arr)
// }

// 3. Array with Implicit Size
// Go can determine the size of the array based on the number of elements you provide,
//  so you don't need to specify the array size explicitly.

// func main() {
// 	arr := [...]int{12, 13, 14, 15, 16, 17}

// 	fmt.Println(arr)
// }

// 4. Array of Strings
// You can have an array of strings, just like any other data type.

// func main() {
// 	var fruits [3]string
// 	fruits[0] = "Apple"
// 	fruits[1] = "Banana"
// 	fruits[2] = "Orange"

// 	fmt.Println("Fruits:", fruits)
// }

// 5. Multi-Dimensional Arrays
// Go supports multi-dimensional arrays, which are arrays of array

// func main() {
// 	var matrix [2][3]int // 2D array: 2 rows and 3 columns
// 	matrix[0][0] = 1
// 	matrix[0][1] = 2
// 	matrix[0][2] = 3
// 	matrix[1][0] = 4
// 	matrix[1][1] = 5
// 	matrix[1][2] = 6

// 	matrix1 := [2][3]int{{1, 2, 3}, {3, 4, 5}}

// 	var matrix2 = [2][3]int{{1, 2, 3}, {3, 4, 5}}

// 	fmt.Println("Matrix:")
// 	fmt.Println(matrix, matrix1, matrix2)
// }

// 6. Array with Fixed Size and Zero Values
// If you don't explicitly initialize the elements of an array,
//  they are automatically initialized to the zero value for their type
//  (e.g., 0 for integers, "" for strings).

// func main() {
// 	var arr [5]int
// 	fmt.Println(arr)
// }

// 7. Array of Structs
// You can also have arrays of structs, which are user-defined types in Go.

// type Person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	var people [2]Person
// 	people[0] = Person{"Alice", 30}
// 	people[1] = Person{"Bob", 25}

// 	fmt.Println("People:", people)
// }

// 8. Slicing an Array
// Although arrays are fixed in size, Go allows you to work with slices,
//  which are dynamic and can be resized. A slice is essentially a view or
//   a reference to a portion of an array.

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slice := arr[1:4] // Slice from index 1 to 3
// 	// slice := arr[:4] // Slice from index 0 to 3

// 	fmt.Println("Original Array:", arr)
// 	fmt.Println("Slice:", slice)
// }

// 9. Array Passed by Value
// In Go, arrays are passed by value to functions.
//  This means a copy of the array is passed,
//  not a reference to the original array.

// func modifyArray(arr [3]int) {
// 	arr[0] = 99 // This modification will not affect the original array
// }

// func main() {
// 	arr := [3]int{1, 2, 3}
// 	modifyArray(arr)
// 	fmt.Println("Array after function call:", arr) // Original array remains unchanged
// }

// 10. Array with Large Sizes
// Go allows creating arrays of a large size,
//  but it's often better to use slices for flexibility.
//   However, here is an example of large arrays:

func main() {
	var largeArr [1000]int
	largeArr[999] = 5000 // Assigning a value to the last element

	fmt.Println("Last element in large array:", largeArr[999])
}
