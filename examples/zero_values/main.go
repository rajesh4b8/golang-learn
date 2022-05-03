// Go program to illustrate the concept of zero value
package main

import "fmt"

// Main Method
func main() {

	// Creating variables
	// of different types
	// var q1 int = 0
	var q1 int
	// var q2 float64 = 0.0
	var q2 float64
	// var q3 bool = false
	var q3 bool
	// var q4 string = ""
	var q4 string
	var q5 []int
	var q6 *int
	var q7 map[int]string

	// Displaying the zero value
	// of the above variables
	fmt.Println("Zero value for integer types: ", q1)
	fmt.Println("Zero value for float64 types: ", q2)
	fmt.Println("Zero value for boolean types: ", q3)
	fmt.Println("Zero value for string types: ", q4)
	fmt.Println("Zero value for slice types: ", q5)
	fmt.Println("Zero value for pointer types: ", q6)
	fmt.Println("Zero value for map types: ", q7)
}
