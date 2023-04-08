package main

import (
	"errors"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(cards))
	}
}

func TestSavingToFile(t *testing.T) {
	cards := newDeck()

	cards.saveToFile("my_cards")

	if _, err := os.Stat("./my_cards"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Failed saving deck to the file system")
	}
}
