// In Go, enums (short for "enumerations") are not directly supported as a special data
//  type like in some other languages. However, Go provides a way to simulate enums using
//   constants and iota, a special Go keyword used for generating a series of constant values.

// Here's an example and explanation of how to use enums in Go:

package main

import "fmt"

// Define an enum for Days of the Week using iota
type Day int

const (
	Sunday Day = iota // iota starts at 0 by default
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// const (
// 	Sunday    = "Sunday"
// 	Monday    = "Monday"
// 	Tuesday   = "Tuesday"
// 	Wednesday = "Wednesday"
// 	Thursday  = "Thursday"
// 	Friday    = "Friday"
// 	Saturday  = "Saturday"
// )

func main() {
	// Printing out the days
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)

	// You can also compare or use the enums
	today := Friday
	if today == Friday {
		fmt.Println("Today is Friday!")
	}
}
