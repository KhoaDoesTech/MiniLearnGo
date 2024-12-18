package tests

import (
	"github.com/KhoaDoesTech/MiniLearnGo/Basic/cards"
	"os"
	"testing"
)

func TestSaveToDeckAndNewDeskFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := cards.NewDeck()
	deck.SaveToFile("_decktesting")

	loadedDeck := cards.NewDeckFromFile("_decktesting")
	if len(loadedDeck) != len(deck) {
		t.Errorf("Loaded deck length is not equal to the original deck length")
	}

	os.Remove("_decktesting")
}
