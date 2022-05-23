package main

import (
	"fmt"
	"time"
)

// 0, 1, 1, 2, 3, 5, 8, 11, etc...
func main() {
	c := make(chan int)
	go fibonacci(10, c)

	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {

	x, y := 0, 1

	for i := 0; i < n; i++ {
		// fmt.Println(x)
		time.Sleep(2 * time.Second)
		c <- x
		x, y = y, x+y
	}

	close(c)
}
