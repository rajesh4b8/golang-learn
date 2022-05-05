package main

import (
	"fmt"
)

type person struct {
	name string
	zip  int
}

func main() {
	// struct value created in memory
	p := person{"Rajesh", 24325}

	fmt.Println("Data", p)

	personPointer := &p
	fmt.Println("Address", personPointer)

	fmt.Println("Data from pointer", *personPointer)

	p.updatePersonValueReceiver("Ramesh")
	fmt.Println("Updated", p)

	// personPointer = &p
	// personPointer.updatePersonPointerReceiver("Suresh")
	// Pointer indirection
	p.updatePersonPointerReceiver("Suresh")
	fmt.Println("updatePersonPointerReceiver", p)
}

// Pass by Value Receiver!
func (p person) updatePersonValueReceiver(name string) {
	p.name = name
}

// Pass by pointer receiver!
func (p *person) updatePersonPointerReceiver(name string) {
	// (*p).name = name
	p.name = name
}

// Reference types -> slice, map
// Value types -> int, string, struct, array
