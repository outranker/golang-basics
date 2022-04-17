package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name
	fmt.Println("bef", name)
	fmt.Println(&namePointer)
	printPointer(namePointer)
	fmt.Println("aft", name)
}

func printPointer(namePointer *string) {
	*namePointer = "some"
	fmt.Println(&namePointer)
}
