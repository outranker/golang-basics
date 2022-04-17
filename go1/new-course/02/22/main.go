package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintf(rw, "This is the home page")
}

func About(rw http.ResponseWriter, r *http.Request) {
	sum := addValues(2,3)
	_,_ = fmt.Fprintf(rw, fmt.Sprintf("this is the about page and 2+3 is %d", sum))
}

func addValues(x, y int) int {
	return x+y;
}

func main() {
	


	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
	// log.Println(a)
}