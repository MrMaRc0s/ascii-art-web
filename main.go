package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii-art-web/printer" // Update this to the correct import path for your project
)

func main() {
	// Serve static files (e.g., CSS) from templates/style directory
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("templates/style"))))

	// Serve templates
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ascii-art", generateAsciiArt)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home page template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}

	// Render the template without ASCII art at first
	tmpl.Execute(w, nil)
}

func generateAsciiArt(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")     // Get the text input from the form
	banner := r.FormValue("banner") // Get the selected banner

	// Ensure the banner is valid
	if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
		http.Error(w, "Invalid banner choice", http.StatusBadRequest)
		return
	}

	// Generate ASCII art using the printer function
	asciiArt, err := printer.GenerateAsciiArt(text, banner)
	if err != nil {
		http.Error(w, "Error generating ASCII art", http.StatusInternalServerError)
		return
	}

	// Render the page again, passing the ASCII art as data
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}

	// Pass the ASCII art to the template
	tmpl.Execute(w, struct{ AsciiArt string }{AsciiArt: asciiArt})
}
