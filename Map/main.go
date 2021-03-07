package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fff",
	}
	printMap(colors)

	var colors2 map[string]string
	printMap(colors2)

	colors3 := make(map[string]string)
	colors3["white"] = "#fff"
	fmt.Println()
	printMap(colors3)

	// delete(colors, "white")

	// fmt.Println(colors, colors2, colors3)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		hex = "testing"
		fmt.Println(color, hex)
	}
}
