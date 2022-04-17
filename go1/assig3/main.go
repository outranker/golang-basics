package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{
	height float64
	base float64
}

type square struct{
	sideLength float64
}

func (t triangle) getArea() float64 {
	a := 0.5 * t.base * t.height
	return a
}
func (s square) getArea() float64 {
	a := s.sideLength * s.sideLength
	return a
}

func printArea(sh shape) {
	v := sh.getArea()
	fmt.Println(v) 

}

func main() {
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}	
	printArea(t)
	printArea(s)
}