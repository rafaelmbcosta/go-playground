package main

import "fmt"

func main() {
	// maps need to be initalized
	// var approved map[int]string

	approved := make(map[int]string)

	approved[123456] = "Maria"
	approved[654321] = "José"

	fmt.Println(approved)

	for number, name := range approved {
		fmt.Printf("%s (NUMBER: %d) \n", name, number)
	}

	fmt.Println(approved[123456], "Printing from key")
	delete(approved, 123456)
	fmt.Println(approved, "After removal")
}
