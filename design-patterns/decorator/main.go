// Structural design patterns: Decorator pattern example
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloEndpoint)
	http.HandleFunc("/logged", logDecorator(helloEndpoint))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func logDecorator(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		fn(w, r)
	}
}
