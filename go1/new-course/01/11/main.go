package main

import "log"
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	// BirthDate time.Time
}

func (l *User) printName() {
	log.Println("printing first name",l.FirstName)
}
func main() {

	user := User{FirstName: "sdsd"}
	var user2 User = user
	user.printName()
	m := `some thing`
	log.Println(m)
	user2.Age = 22;
	log.Println(user2)
	
	log.Println(user)
}