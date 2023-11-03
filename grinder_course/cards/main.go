package main

import "fmt"

func main() {
	cards := newDeck() // Nome disso é SLICE!!
	cards.print()

	hand, remaining_cards := deal(cards, 5)
	fmt.Println("Deal (Hand) ", hand)
	fmt.Println("Remaining Cards ", remaining_cards)
	// fmt.Println(cards.saveToFile("saved deck"))
	fmt.Println(newDeckFromFile("saved deck"))
}
