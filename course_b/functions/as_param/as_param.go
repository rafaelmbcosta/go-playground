package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func sum(a, b int) int {
	return a + b
}

func exec(function func(int, int) int, p1, p2 int) int {
	return function(p1, p2)
}

func main() {
	fmt.Println("SOMA =", sum(3, 4))
	result := exec(multiply, 3, 4)
	fmt.Println(result)
}
