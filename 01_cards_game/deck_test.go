package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}
	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deck[0])
	}
	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs, but got %v", deck[len(deck)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	filename := "_decktesting.txt"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	_, err := os.Stat(filename)
	if err != nil {
		t.Errorf("No file found with the name %v", filename)
		return
	}
}

func TestNewDeckFromFile(t *testing.T) {
	filename := "_decktesting.txt"
	deck, err := newDeckFromFile(filename)

	if err != nil {
		t.Errorf("No file found with the name %v", filename)
		return
	}
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
		return
	}
	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deck[0])
	}
	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs, but got %v", deck[len(deck)-1])
	}
}
