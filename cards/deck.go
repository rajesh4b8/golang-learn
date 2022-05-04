package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// Function with Receiver argument, as called Method on type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Function with Return argument
func createNewDeck() deck {
	suits := []string{"Spades", "Daimonds", "Clubs", "Hearts"}
	values := []string{"Ace", "Two", "Three", "Four"} // should be 13 but only definining 4 for simplicity

	newDeck := deck{}

	// for loop by using range
	for _, suit := range suits {

		// another way of doing for loop
		for j := 0; j < len(values); j++ {
			newDeck = append(newDeck, values[j]+" of "+suit)
		}
	}

	return newDeck
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) writeToFile(fileName string) {
	str := d.toString()
	err := ioutil.WriteFile(fileName, []byte(str), 0666)
	if err != nil {
		println("Error saving the file:", fileName, err.Error())
		os.Exit(1)
	}
}

func readFromFile(fileName string) deck {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		println("Error reading the file:", fileName, err.Error())
		os.Exit(1)
	}

	str := string(content)
	slice := strings.Split(str, ",")

	// type conversion to deck
	return deck(slice)
}

func (d deck) shuffle() {
	// create new source from a variable seed0
	source := rand.NewSource(time.Now().UnixNano())
	// create a random from source
	r := rand.New(source)
	for i := range d {
		// give random number between 0 and len - 1
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
