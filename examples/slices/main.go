package main

import "fmt"

func main() {
	cards := []string{"Zero", "One", "Two", "Three", "Four"}

	// print all values
	fmt.Println(cards)

	// first two
	fmt.Println(cards[:2])

	// everything after two
	fmt.Println(cards[2:])

	// print only one and two
	fmt.Println(cards[1:3])
}
