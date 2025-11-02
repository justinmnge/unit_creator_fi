package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files (CSS, images, etc) from the static directory
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Handle the homepage
	mux.HandleFunc("/", handleHome)

	// Handle the portal page
	mux.HandleFunc("/login", handleLogin)

	// Handle the trailer page
	mux.HandleFunc("/trailer", handleTrailer)

	fmt.Println("Server listening to :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html")
}

func handleTrailer(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "trailer.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("static/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
