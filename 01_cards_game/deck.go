package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	cards := deck{}
	cardsSuit := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardsSuit {
		for _, value := range cardsValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	file, err := os.ReadFile(filename)
	var d deck

	sliceStrings := strings.Split(string(file), ",")
	d = deck(sliceStrings)
	return d, err
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	rand.New(src)
	for i := range d {
		newIndex := rand.Intn(len(d) - 1)
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}
