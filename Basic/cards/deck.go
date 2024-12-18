package cards

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []string

func NewDeck() deck {
	cardSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Preallocate memory for the slice to optimize performance
	// Measure performance with 10000 times: 24.4312ms
	cards := make(deck, 0, len(cardSuits)*len(cardValues))

	// Measure performance with 10000 times: 32.9332ms
	//cards := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Printf("%d: %s\n", i+1, card)
	}
}

func Deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] // Swap
	}
}
