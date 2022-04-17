package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	var abc string
	dir, file = path.Split("css/")
	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
	fmt.Println(abc)
}