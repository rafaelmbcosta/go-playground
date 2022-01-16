package main

import "fmt"

func main() {
	// initializing map
	funcsAndSalary := map[string]float64{
		"John":   1122.0,
		"Fabio":  14400,
		"Thomas": 1299, // final comma is needed in this layout
	}

	fmt.Println(funcsAndSalary)

	for name, salary := range funcsAndSalary {
		fmt.Println(name, salary)
	}
}
