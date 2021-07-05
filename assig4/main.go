package main

import (
	"fmt"
	"log"
	"os"
)

// tasks

// 1. check out os.Args - done
// 2. learn how to print agrs/input by user and save that to a variable - done
// 3. open a file with os.Open and print everything that's written in it - half done
// 4. also check out os pkg
// 5. interfaces, i guess?

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error occured while opening file: ",err)
		os.Exit(1)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	
}