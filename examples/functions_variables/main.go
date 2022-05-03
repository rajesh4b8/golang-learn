package main

import "fmt"

func main() {

	func() {
		fmt.Println("Print now!!!")
	}()

	// defer would defer the execution of the function to the end of current function
	// defer would execute the function before returning
	defer func() {
		fmt.Println("Defer - Print!!!")
	}()

	addition := func(i, j int) int {
		return i + j
	}

	substraction := func(i, j int) int {
		return i - j
	}

	fmt.Println("2 + 4", addition(2, 4))
	fmt.Println("2 - 4", substraction(2, 4))
	fmt.Println("Op 2 + 4", doAurth(addition, 2, 4))
	fmt.Println("Op 2 - 4", doAurth(substraction, 2, 4))

}

// func add(i, j int) int {
// 	return i + j
// }

func doAurth(op func(int, int) int, i int, j int) int {
	return op(i, j)
}
