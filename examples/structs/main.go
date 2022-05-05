package main

import "fmt"

// struct definition
type person struct {
	firstName string
	lastName  string
	age       int
	addr      address
}

type address struct {
	city    string
	zipcode int
}

func main() {
	// var p person = person{"Rajesh", "Reddy", 12}
	// p := person{"Rajesh", "Reddy", 12}
	// p := person{lastName: "Reddy", firstName: "Rajesh", age: 12}
	p := person{lastName: "Reddy", firstName: "Rajesh"}

	// when we don't define any values.... it's going to assign the zero values
	// p := person{}

	// updating the values
	p.age = 12

	// fmt.Println(p)
	fmt.Printf("%+v \n", p)

	np := person{
		lastName:  "Reddy",
		firstName: "Rajesh",
		age:       12,
		addr: address{
			city: "Phoenix",
			// zipcode: 85248,
		},
	}
	np.addr.zipcode = 85248

	// np.updateAddress("Dallas", 45645)
	np.addr.updateAddress("Dallas", 243534)

	fmt.Printf("%+v %v %T  \n", np, np, np)
}

// func (p person) updateAddress(city string, zip int) {
// 	p.addr.city = city
// 	p.addr.zipcode = zip
// 	fmt.Printf("Inside updateAddress:: %+v \n", p)
// }

func (a *address) updateAddress(city string, zip int) {
	a.city = city
	a.zipcode = zip
	fmt.Printf("Inside updateAddress:: %+v \n", a)
}
