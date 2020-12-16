package main

import "fmt"

func main() {
	var stringVariable1 string = "Variable 1"
	stringVariable2 := "Variable 2"

	fmt.Println(stringVariable1)
	fmt.Println(stringVariable2)

	var (
		variable3 string = "v3"
		variable4 string = "v4"
	)

	// sort of swap
	variable4, variable3 = variable3, variable4

	fmt.Println(variable3)
	fmt.Println(variable4)

	var5, var6 := "v5", "v6"

	fmt.Println(var5)
	fmt.Println(var6)

	const const1 string = "Constant"
	fmt.Println(const1)

}
