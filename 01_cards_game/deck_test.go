package main

import "testing"

//t.Errorf("Expected deck length of 16, but got %v", len(d))
func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}
}

func TestNewDeckFromFile(t *testing.T) {
	filename := "my_cards.txt"
	deck, err := newDeckFromFile("my_cards.txt")
	if err != nil {
		t.Errorf("No file found with the name %s", filename)
		return
	}
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
		return
	}
}
