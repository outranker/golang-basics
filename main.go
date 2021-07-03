package main

func main() {
	cards := newDeck()
	cards.shuffle()

	cards.print()
	numbers := []int{0,1,2,3,4,5,6,7,8,9,10}
	isEven(numbers)

}

