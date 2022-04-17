package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#67b762",
		"white": "#ffffff",
	}
	delete(colors, "red")

	fmt.Println(colors)
}
