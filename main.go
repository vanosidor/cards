package main

import "fmt"

func main() {
	cards := newDeck()

	cards = cards[:5]
	fmt.Println(cards.toString())

	cards.shuffle()
	fmt.Println(cards.toString())

	newDeck := newDeckFromFile("deck1.txt")

	newDeck.print()
}
