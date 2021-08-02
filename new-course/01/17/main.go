package main

import (
	"log"

	"github.com/outranker/myniceprogram/helpers"
)

const np = 1000

func cv(c chan int) {
	rn := helpers.RandomNumber(np)

	c <- rn
}


func main() {
	c := make(chan int)

	defer close(c)
	go cv(c)

	num := <-c
	log.Println(num)
}