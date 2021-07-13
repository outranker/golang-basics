package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
	fmt.Println(rand.Intn(55533321212))
	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 8
	password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
	fmt.Print("minNum: ", minNum, "\nminSpecialChar: ", minSpecialChar, "\nminUpperCase: ", minUpperCase,
		"\npasswordLength: ", passwordLength, "\npassword: ", password)
	// fmt.Print("\nallCharSet: ",allCharSet)

	fmt.Println(password)
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder
	fmt.Println("testing how len works for strings", string(specialCharSet[1]))
	fmt.Println("AAaaAAAAAA:: ", string(specialCharSet[3]))
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		println("random", random)
		password.WriteString(string(specialCharSet[random]))
		fmt.Println("this is all after: ", password)
	}
	return ""
}
