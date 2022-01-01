package main

import (
	"fmt"
)

func arrays() {
	var notes [3]float64

	notes[0], notes[1], notes[2] = 1.5, 2.5, 3.8

	fmt.Println(notes)

	total := 0.0
	for i := 0; i < len(notes); i++ {
		total += notes[i]
	}

	average := total / float64(len(notes))
	fmt.Printf("Average is %.2f\n", average)
}

func for_range() {
	numbers := [...]int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		fmt.Printf("%d) %d\n", i+1, number)
	}

	for _, number := range numbers {
		fmt.Printf("%d\n", number)
	}
}

func main() {
	// arrays()
	// fmt.Println("oi")
	for_range()
}
