package main

import "log"

func main() {
	isTrue := false
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}
	
	
	cat := "cat"

	if cat == "cat" {
		log.Println("cat is", cat) 
	} else  {
		log.Println("cat is not", cat)
	}

	n := 100
	t := true
	if n > 99 && t {
		log.Println("n is greater than 100 and t is true")
	} else if n < 100 && t { 
		log.Println("1")
	} else if n == 100 || t {
		log.Println("2")
	} else if n > 1000 && t  == false {
		log.Println("3")
	}

c := "dog"
	switch c {
	case "cat":
		log.Println("this is c", c)
	case "dog":
		log.Println("this is c", c)
	default: 
	log.Println("this is c from default", c)
	}
}