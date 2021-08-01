package main

import "log"

type User struct {
	FirstName string
	LastName string
}
func main() {
	// this below not working no?
// var myOtherMap map[string]string
// myOtherMap["hello"] = "world"
// log.Println("first one:", myOtherMap["hello"])

myMap := make(map[string]string)
myMap["what's up,"] = "buddy?"
log.Println("second one:", myMap["what's up,"])

m := make(map[string]int)
m["boi"]= 21
log.Println(m["boi"])

l := make(map[string]User)
me := User{
	FirstName: "yoboi",
	LastName: "sup?",
}
l["maboi"] = me
log.Println(l["maboi"])

// slices

var mySlice []string
mySlice = append(mySlice, "Trevor")
mySlice = append(mySlice, "John")
mySlice = append(mySlice, "Mary")
log.Println(mySlice)

var newSl []int
newSl = append(newSl,1)
newSl = append(newSl,1)
newSl = append(newSl,1)
log.Println(newSl)

n := []int{3,4,5,6}
log.Println("d&i", n)

names := []string{"one", "two", "three"}

log.Println(names)
}