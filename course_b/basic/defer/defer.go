package main

import "fmt"

func getValue(number int) int {
	defer fmt.Println("end")

	if number > 5000 {
		fmt.Println("higher...")
		return 5000
	}
	fmt.Println("lower...")
	return number
}

func main() {
	fmt.Println(getValue(7000))
	fmt.Println(getValue(3000))
}
