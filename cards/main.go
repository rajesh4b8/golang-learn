package main

func main() {
	newDeck := createNewDeck()

	// hand, remainingCards := deal(newDeck, 5)
	// fmt.Println("hand")
	// hand.print()
	// fmt.Println("remaining ", remainingCards.toString())

	// newDeck.writeToFile("current_deck.txt")

	// deckFromFile := readFromFile("current_deck.txt")

	newDeck.shuffle()
	newDeck.print()
}
