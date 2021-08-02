package main

import "log"

func main() {

	for i := 0; i <= 0; i++ {
		log.Println(i)
	}

	animals := []string{"dog","fish","horse","cow"}

	for _, animal := range animals {
		log.Println(animal)
	}

	a := make(map[string]string)

	a["n1"]= "dog"
	a["n2"]= "cat"
	a["n3"]= "cow"

	for i, n := range a {
		log.Println(i, n)
	}

	l := "some random string"

	for a,b := range l {
		log.Println(a,b)
	}

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	user := User{FirstName: "J", LastName: "F", Email: "e@d.c", Age: 26}

	users:=make([]User, 0)
	users = append(users, user)

	for a,b := range users {
		log.Println(a,b)
	}
}
