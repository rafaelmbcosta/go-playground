package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// create a new type of deck, which is a slice of strings

type deck []string

// function with receiver, every instance of type deck will have access
// to the print function.
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Swords"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	err := ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
	return err
}
