package main

import (
	"fmt"
)

func main() {
	// var msg string = "Hi there"
	msg := "Hi there"
	fmt.Println(msg)

	msg = "10"

	myBool := true
	myInt := 20
	var myFloat float32 = 23.65

	fmt.Printf("type is %T, value is %v \n", msg, msg)
	fmt.Printf("type is %T, value is %v \n", myBool, myBool)
	fmt.Printf("type is %T, value is %v \n", myInt, myInt)
	fmt.Printf("type is %T, value is %v \n", myFloat, myFloat)

	// buit in types

	/*
		bool, string,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		rune, byte,
		float32, float64,
		complex64, complex128
	*/

	status, _ := printMsg("Something")
	fmt.Println("Status", status)
}

func printMsg(msg string) (bool, string) {
	fmt.Println(msg)
	return true, msg
}
