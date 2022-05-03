package main

import "fmt"

func main() {
	var brand string
	brand = "Google"
	fmt.Printf("%q\n", brand)
	fmt.Printf("%s\n", brand)
	fmt.Println("hello\\")

	var some = 16.4
	fmt.Printf("this %T\n", some)
	fmt.Printf("this is %v, baby! i said %[1]v!\n", brand)
}