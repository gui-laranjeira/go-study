package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors)

	colors["white"] = "#ffffff"
	fmt.Println(colors)

	delete(colors, "white")
	fmt.Println(colors)
	itareteMap(colors)
}

func itareteMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key + ": " + value)
	}
}
