// Lesson 1: Prep
package main

import (
	"net/http"
	"log"
	"fmt"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "<h1>Hello! My name is %s</h1>", name)
}


func main() {
	http.HandleFunc("/", greetings)
	log.Println("http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}