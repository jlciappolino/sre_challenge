package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Estas en users, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	log.Fatal(http.ListenAndServe(":8181", nil))
}
