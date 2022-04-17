package main

import (
	"log"
)

var s = "seven"



func main() {
var s2 = "six"
log.Println("s is", s)
log.Println("s2 is", s2)

say("xxx")


}

func say(s3 string) (string, string) {
	log.Println("s from say func is", s)
	return s3, "World"
}