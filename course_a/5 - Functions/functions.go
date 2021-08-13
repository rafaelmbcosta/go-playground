package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculations(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	minus := n1 - n2
	return sum, minus
}

func main() {
	fmt.Println(sum(2, 7))

	// declaring like 'javascript arrow functions'
	var f = func(text string) string {
		fmt.Println(text)
		return text
	}

	f("Opa!")

	// resultSum, resultSub := calculations(10, 12)

	// ignores second result
	resultSum, _ := calculations(10, 12)
	fmt.Println(resultSum)

	// ignores first result
	_, resultSub := calculations(10, 12)
	fmt.Println(resultSub)

}
