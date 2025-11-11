package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files (CSS, images, etc) from the static directory
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Handle the homepage
	mux.HandleFunc("/", handleHome)

	// Handle the portal page
	mux.HandleFunc("/login", handleLogin)

	// Handle the trailer page
	mux.HandleFunc("/trailer", handleTrailer)

	// Handle the contact page
	mux.HandleFunc("/contact", handleContact)

	// Hanlde the about page
	mux.HandleFunc("/about", handleAbout)

	fmt.Println("Server listening to :8080")
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed: %v\n", err)
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

func handleContact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("static/" + tmpl)
	if err != nil {
		log.Printf("Error parsing template %s: %v", tmpl, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template %s: %v", tmpl, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
