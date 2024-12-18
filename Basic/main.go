package main

import (
	"fmt"
	"github.com/KhoaDoesTech/MiniLearnGo/Basic/cards"
	"os"
	"time"
)

const age int = 21

// Variables can't be declared and initialized using the short declaration operator (`:=`) outside of function
// age := 20

func init() {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		fmt.Println("APP_NAME is not set!")
	} else {
		fmt.Println("App Name:", appName)
	}
	fmt.Printf("Khoa is %d years old.\n", age)
}

func main() {
	fullCards := cards.NewDeck()
	fullCards.Shuffle()
	fullCards.Print()

	//fullCards := cards.NewDeckFromFile("my_cards.txt")
	//fullCards.Print()

	//measureExecutionTime("NewDeck", func() { cards.NewDeck() })
}

func measureExecutionTime(name string, action func()) {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		action() // Execute the provided function
	}
	elapsed := time.Since(start)
	fmt.Printf("\nExecution time of %s: %s\n", name, elapsed)
}
