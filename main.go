package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// fmt.Println(card)
	// card = "Five of Diamonds"
	// fmt.Println(card)
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	fmt.Println(cards)

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
