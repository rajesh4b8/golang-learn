package main

import "fmt"

func main() {
	// var card string
	// card = "Four of Spades"

	// var card string = "Ace of diamonds"
	card := "Ace of diamonds"
	card = "Four of Spades"

	// var cards []string = []string{"Ace of Spades", card}
	cards := []string{"Ace of Spades", card, newCard()}

	print(cards)
}

func newCard() string {
	return "Six of Clubs"
}

func print(cards []string) {
	// var card string
	// var i int
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
