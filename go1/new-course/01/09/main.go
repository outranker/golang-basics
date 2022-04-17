package main

import "fmt"

func main() {
	s := "green"

	fmt.Println("before: ", s)
	c(&s)
	fmt.Println("after: ", s)
}

func c(s *string) {
	*s = "red"
}