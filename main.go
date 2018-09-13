package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.print()

	hand, remaingDeck := deal(cards, 5)
	hand.print()
	remaingDeck.print()

	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
