package main

import "fmt"

func main() {
	// Map declaration
	// var shapes map[string]int

	// Initializaing the map using make()
	// var shapes map[string]int = make(map[string]int)

	// creating the empty map
	shapes := make(map[string]int)

	shapes["Triangle"] = 3
	shapes["Square"] = 4
	shapes["Rectangle"] = 4

	colors := map[string]string{
		"Red":  "#124235",
		"Blue": "#75474",
	}

	colors["Red"] = "#98733"
	updateColorCode(colors)

	printMap(colors)
}

func updateColorCode(c map[string]string) {
	c["Black"] = "#000000"
}

func printMap(c map[string]string) {
	for color, code := range c {
		fmt.Println(color, "::", code)
	}
}
