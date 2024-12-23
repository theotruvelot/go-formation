package main

func main() {
	//Plusieurs façon de déclarer des variables
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	//card := newCard()
	//fmt.Println(card)

	//utilisation du type deck
	cards := newDeck()

	cards.print()
}
