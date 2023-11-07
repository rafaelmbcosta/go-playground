package main

import "fmt"

type numbers []int

func newNumbers() numbers {
	n := numbers{}

	for i := 0; i <= 10; i++ {
		n = append(n, i)
	}

	return n
}

func (n numbers) print() {
	for _, element := range n {
		if element%2 == 0 {
			fmt.Printf("Number %v is even\n", element)
		} else {
			fmt.Printf("Number %v is odd\n", element)
		}
	}
}

func (n numbers) oddNumbers(index int) bool {
	if n[index]%2 == 0 {
		return false
	} else {
		return true
	}
}

func (n numbers) evenNumbers(index int) bool {
	return !n.oddNumbers(index)
}
