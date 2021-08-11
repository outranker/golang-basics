package main

import (
	"fmt"
	"log"
	"net/http"
)

// task 1
// - print path param
// - if even redirect else no redirect
// task 2
// - open, parse, read yaml file
// - save yaml to array of structs and loop through it once


func main() {
	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://localhost:3005/hello", 307)

	mux.Handle("/abc", rh)
	mux.Handle("/hello", A)
	log.Println("Listening at 3005")
	http.ListenAndServe(":3005", mux)
	
	
}
func A(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}
// filePath := flag.String("file", "redirect.yaml", "file path name for the yaml file")
// flag.Parse()
// ParseYaml(filePath)

// lines := make(map[string]string)
// for _, m := range yml{
// 	lines[m["path"]] = m["url"]
// }

// fmt.Println(lines)

// helloHandler := func(w http.ResponseWriter, req *http.Request){
// 	fmt.Println("this is url path param", req.URL.Path)
// }
// http.HandleFunc("/:value", helloHandler)
// _ = http.ListenAndServe(":3005", nil )
// http.ServeMux

// func ParseYaml(filepath *string) []map[string]string {
	
// 	file, err := ioutil.ReadFile(*filepath)
// 	if err != nil {
// 		fmt.Println("Error opening file: ", err)
// 	}
// 	yml := []map[string]string{}
// 	err = yaml.Unmarshal(file, &yml)
// 	if err != nil {
// 		fmt.Println("error unmarshalling yaml: ", err)
// 	}
// 	return yml
// }
