package main

func main() {
	//Plusieurs façon de déclarer des variables
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	//card := newCard()
	//fmt.Println(card)

	//utilisation du type deck
	cards := deck{"Ace of Diamonds", newCard()}
	//ajout dans un slice
	cards = append(cards, "Six of Spades")
	//boucle
	cards.print()
}

// Type of return = string
func newCard() string {
	return "Five of Diamonds"
}
