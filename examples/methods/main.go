package main

import "fmt"

// Starts with upper case letter... It's visible outside the package
type Deck []string

// Starts with lower case letter... It's only visible in the same package
type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Rajesh", age: 12}
	p.Print()
	fmt.Println(p.changeName("Ramesh"))

	cards := Deck{"Ace", "Two"}
	cards.Print()
}

func (d Deck) Print() {
	fmt.Println(d)
}

func (p person) Print() {
	fmt.Println(p)
}

func (p person) changeName(newName string) person {
	p.name = newName

	return p
}
