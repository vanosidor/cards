package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	cards := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards

}

func deal(cards deck, handsize int) (deck, deck) {
	return cards[:handsize], cards[handsize:]
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile("deck.txt", []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bytes, error := os.ReadFile(filename)

	if error != nil {
		os.Exit(1)
	}

	deck := strings.Split(string(bytes), ",")
	return deck
}

func (d deck) shuffle() deck {

	for index := range d {
		randomIndex := rand.Intn(len(d))

		// TODO swap
		d[index], d[randomIndex] = d[randomIndex], d[index]
	}

	return d
}
