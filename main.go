package main

import "fmt"

func main() {
	//Plusieurs façon de déclarer des variables
	//var card string = "Ace of Spades"
	card := newCard()
	fmt.Println(card)
}

// Type of return = string
func newCard() string {
	return "Five of Diamonds"
}
