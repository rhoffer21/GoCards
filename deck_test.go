package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 instead got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, instead it is %v", d[0])
	}

	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected the last card to be King of Hearts, instead it is %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52 instead got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
