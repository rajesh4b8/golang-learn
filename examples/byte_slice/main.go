package main

import "fmt"

func main() {
	// byte -> 8 bits
	// bit -> 0 or 1

	// byte range -> 00000000 to 11111111
	// byte in decimal -> 0 to 255

	// string to byte slice
	str := "Hi there!"
	content := []byte(str)
	fmt.Println(str, content)

	// ASCII representation of string -> [72 105 32 116 104 101 114 101 33]

	fmt.Println("[]byte to string:", string(content))

}
