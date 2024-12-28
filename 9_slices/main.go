package main

import "fmt"

func main() {
	// 	1.Empty Slice
	// An empty slice is a slice with no elements but can hold elements as it is expanded.
	// var emptySlice []int    // Declaring an empty slice
	// fmt.Println(emptySlice) // Output: []

	// 	2.Slice with Literal Values
	// A slice can be initialized with a list of values.
	// slice := []int{1, 2, 3, 4}
	// fmt.Println(slice) // Output: [1 2 3 4]

	// 	3.Slice from an Array
	// You can create a slice by referencing a segment of an array.
	//  A slice does not copy the data from the array; it refers to the original array's underlying array.

	// arr := [5]int{10, 20, 30, 40, 50}
	// slice := arr[1:4]  // Creates a slice from index 1 to 3 (not including 4)
	// fmt.Println(slice) // Output: [20 30 40]

	// 4.Slice with make() Function
	// You can create a slice with a specific length and capacity using the make() function.

	// slice := make([]int, 3, 5) // Length 3, Capacity 5
	// fmt.Println(slice)         // Output: [0 0 0]
	// fmt.Println(len(slice))    // Output: 3 (length of the slice)
	// fmt.Println(cap(slice))    // Output: 5 (capacity of the slice)

	// 	5.Slicing a Slice
	// You can slice a slice to obtain a subset of elements.
	// slice := []int{1, 2, 3, 4, 5}
	// subSlice := slice[1:4]  // Slicing from index 1 to 3 (not including 4)
	// fmt.Println(subSlice)  // Output: [2 3 4]

	// 6.Appending to a Slice
	// You can use the append() function to add elements to a slice.

	// slice := []int{1, 2, 3}
	// slice = append(slice, 4, 5) // Appending multiple elements
	// fmt.Println(slice)          // Output: [1 2 3 4 5]

	// 	7.Slice of a Slice
	// A slice can be sliced further to create another slice.
	// slice := []int{1, 2, 3, 4, 5}
	// subSlice := slice[1:4]     // Slice from index 1 to 3
	// subSlice2 := subSlice[0:2] // Slice from index 0 to 1 (within subSlice)
	// fmt.Println(subSlice2)     // Output: [2 3]

	// 	8.Nil Slice
	// A slice that has not been initialized or has been explicitly set to nil.

	// var slice []int           // nil slice
	// fmt.Println(slice)        // Output: []
	// fmt.Println(slice == nil) // Output: true

	// 	9.Slice with copy() Function
	// You can copy elements from one slice to another using the copy() function.
	source := []int{10, 20, 30}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Println(destination) // Output: [10 20 30]

}
