package tests

import (
	"github.com/KhoaDoesTech/MiniLearnGo/Basic/cards"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := cards.NewDeck()
	if len(deck) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(deck))
	}
}

func TestShuffle(t *testing.T) {
	deck := cards.NewDeck()
	originalDeck := make([]string, len(deck))
	copy(originalDeck, deck)

	deck.Shuffle()

	if len(deck) != 52 {
		t.Errorf("Expected shuffled deck length to be 52, but got %d", len(deck))
	}

	if deck[0] == originalDeck[0] && deck[1] == originalDeck[1] {
		t.Error("Shuffle did not change the order of the deck")
	}
}
