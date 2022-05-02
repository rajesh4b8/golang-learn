package main

import "fmt"

func main() {
	newDeck := createNewDeck()

	hand, remainingCards := deal(newDeck, 5)
	fmt.Println("hand")

	hand.print()

	fmt.Println("remaining ", remainingCards.toString())
}
