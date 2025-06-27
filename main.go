package main

import (
	"milk_and_muscles/views/home"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	home := home.Index()
	fs := http.FileServer(http.Dir("public"))

	// serve files in public directory
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.Handle("/", templ.Handler(home))
	http.ListenAndServe(":8080", nil)
}
