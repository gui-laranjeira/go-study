package main

import "log"

func main() {
	// cards := newDeck()

	// err := cards.saveToFile("my_cards.txt")
	// if err != nil {
	// 	log.Panic(err)
	// }
	newDeck, err := newDeckFromFile("my_cards.txt")
	if err != nil {
		log.Flags()
	}
	newDeck.print()
	newDeck.shuffle()
	println("\n\n Shuffled cards:\n")
	newDeck.print()
}
