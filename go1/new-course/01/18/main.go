package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}

func main() {
	
	myJson := `[
		{
			"first_name":"Clark",
			"Last_name":"Kent",
			"hair_color":"black",
			"has_dog": true
		},{
			"first_name":"Bruce",
			"Last_name":"Wayne",
			"hair_color":"black",
			"has_dog": false
		}
	]`

	var unmarshalled []Person
// fmt.Printf("%v",[]byte(myJson))
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled[0])

	var s []Person
	var m1 Person

	m1.FirstName = "Wally"
	m1.HairColor = "West"
	m1.LastName = "red"
	m1.HasDog = false
s = append(s, m1)
	var m2 Person

	m2.FirstName = "Diana"
	m2.HairColor = "Prince"
	m2.LastName = "black"
	m2.HasDog = false
s = append(s, m2)

newJson, err := json.MarshalIndent(s, "", "   ")
if err != nil {
	log.Println("error while marshalling struct into json")
}
log.Println("this is what comes back from marshalling",string(newJson))
}