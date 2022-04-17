package main

import (
	"fmt"
	"net/http"
)

func main() {
	

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(rw, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})


	_ = http.ListenAndServe(":3030", nil)
}