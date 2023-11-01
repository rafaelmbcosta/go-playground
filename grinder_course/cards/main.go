package main

import "fmt"

func main() {
	cards := []string{"A", "B"} // Nome disso é SLICE!!
	cards = append(cards, "C")
	fmt.Println(cards)

	for _, card := range cards {
		fmt.Println(card)
	}
}
