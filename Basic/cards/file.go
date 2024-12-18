package cards

import (
	"fmt"
	"os"
	"strings"
)

// Receiver Functions
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0644)
}

func NewDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bytes), ",")

	return deck(s)
}
