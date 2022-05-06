package main

import "fmt"

type vehicle interface {
	getWheels() string
}

type bike struct{}

type car struct{}

// implementing vehicle interface for bike
func (b bike) getWheels() string {
	return "2 wheels"
}

// implementing vehicle interface for car
func (c car) getWheels() string {
	return "4 wheels"
}

func main() {
	bike := bike{}
	car := car{}

	printSpecs(bike)
	printSpecs(car)

	// passing Stringers to Println() so String() will be used to print.
	fmt.Println(bike, car)

}

func printSpecs(v vehicle) {
	fmt.Println(v.getWheels())
}

// implementing Stringer interface for bike
func (b bike) String() string {
	return "Specs: " + b.getWheels()
}

// implementing Stringer interface for car
func (c car) String() string {
	return "Specs: " + c.getWheels()
}
