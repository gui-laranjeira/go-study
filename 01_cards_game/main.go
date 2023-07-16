package main

import (
	"log"
)

func main() {
	filename := "my_cards.txt"
	cards := newDeck()

	err := cards.saveToFile(filename)
	if err != nil {
		log.Panic(err)
	}

	newDeck, err := newDeckFromFile(filename)
	if err != nil {
		log.Flags()
	}

	newDeck.print()

	newDeck.shuffle()
	println("\n\n Shuffled cards:\n")

	newDeck.print()
}
