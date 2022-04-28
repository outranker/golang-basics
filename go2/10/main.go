package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "carl"
	name2 := "İnanç"
	fmt.Println(len(name))
	fmt.Println(len(name2))
	realCount := utf8.RuneCountInString(name2)
	fmt.Println("realCount: ", realCount)
}