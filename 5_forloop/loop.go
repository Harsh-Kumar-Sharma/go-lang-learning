package main

import "fmt"

//in go lang we have not while loop construct

func main() {
	//while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)

	// 	i++
	// }

	//infinite loop

	// for {
	// 	fmt.Println("testing")
	// }

	//classic loop

	// for i := 0; i <= 3; i++ {
	// 	// break
	// 	if i == 2 {
	// 		continue // miss the 2 value
	// 	}
	// 	fmt.Println(i)
	// }

	//range loop

	for i := range 3 {
		fmt.Println(i)
	}

}
