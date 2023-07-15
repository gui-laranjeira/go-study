package main

import "log"

func main() {
	cards := newDeck()

	hand, base := deal(cards, 5)

	hand.print()
	println("\nSTOP\n")
	base.print()

	println([]byte(hand.toString()))
	err := hand.saveToFile("my_cards.txt")
	if err != nil {
		log.Fatal(err)
	}
}
