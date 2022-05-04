package main

import (
	"os"
	"testing"
)

const (
	testFileName = "_decktesting"
)

func TestNewDeck(t *testing.T) {
	newDeck := createNewDeck()

	if len(newDeck) != 16 {
		t.Errorf("Expected 16 cards but got %d", len(newDeck))
	}

	if newDeck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", newDeck[0])
	}
	if newDeck[15] != "Four of Hearts" {
		t.Errorf("Expected Four of Hearts but got %v", newDeck[15])
	}
}

func TestSaveToDeckAndReadFromDeck(t *testing.T) {
	// make sure to remove the file if any before testing
	os.Remove(testFileName)

	deck := createNewDeck()
	deck.writeToFile(testFileName)
	loadedDeck := readFromFile(testFileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards but got %d", len(loadedDeck))
	}

	os.Remove(testFileName)
}
