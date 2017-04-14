package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/about", handlerAbout)

	http.ListenAndServe(":9010", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\nAlan Here")
}
func handlerAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "More About Me")
}
