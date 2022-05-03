package main

import "fmt"

func main() {
	simpleFunction()

	// var name string, age int = "Age", 12
	name, age := "Age", 12
	functionWithParams(name, age)

	// var i, j int = 5, 10
	i, j := 5, 10
	functionWithParamsOfSameType(i, j)

	fmt.Println(functionWithReturnValue())

	x, y := functionReturnSameValues(2, 4)
	fmt.Println(x, y)
	fmt.Println(functionWithMultiReturns())
	fmt.Println(incrementNumbers(2, 3))
}

func simpleFunction() {
	fmt.Println("Print something")
}

func functionWithParams(param1 string, param2 int) {
	fmt.Println(param1, param2)
}

func functionWithParamsOfSameType(param1, param2 int) {
	fmt.Println(param1, param2)
}

func functionWithReturnValue() int {
	return 2325
}

func functionWithMultiReturns() (string, int) {
	return "Age", 12
}

func functionReturnSameValues(param1, param2 int) (int, int) {
	return param1, param2
}

func incrementNumbers(i, j int) (int, int) {
	i++
	j++
	return i, j
}
