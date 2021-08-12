package main

import (
	"fmt"
	"net/http"
	urlshort "url-shortener/handler"
)

func main() {

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/main":    "https://github.com/gophercises/urlshort/blob/master/main/main.go",
		"/handler": "https://github.com/gophercises/urlshort/blob/master/handler.go",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	mux.HandleFunc("/some-route", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r, "Hello, world!")
		var a int = 1
		fmt.Printf("GET SOME: %x", a)
	})

	fmt.Println("starting the server on :3005")
	http.ListenAndServe(":3005", mapHandler)

}

func defaultMux() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
}
