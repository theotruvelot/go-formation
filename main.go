package main

func main() {
	//Plusieurs façon de déclarer des variables
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	//card := newCard()
	//fmt.Println(card)

	////création d'une slice
	//cards := []string{"Ace of Diamonds", newCard()}
	////ajout dans un slice
	//cards = append(cards, "Six of Spades")
	////boucle
	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}

}

// Type of return = string
func newCard() string {
	return "Five of Diamonds"
}
