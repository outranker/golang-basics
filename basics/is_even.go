package main

import "fmt"


func isEven(number []int) {
	for _, num := range number {
		if num % 2 > 0 {
			fmt.Println("this number is odd: ", num)
		} 
		if num % 2 == 0 {
			fmt.Println("this number is even: ", num)
		}
	}
}