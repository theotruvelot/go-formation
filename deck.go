package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// create receiver function for type deck
// any variable with type deck have access to print func
// the receiver variable need to be an abbreviation of the type (here d)
// but we can make dec / de... BUT not cards... (you can but convention)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
