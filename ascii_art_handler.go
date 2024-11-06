package main

import (
	"net/http"
	"text/template"
)

// ArtData holds data to be passed to templates (input, generated ASCII art, banner, errors)
type ArtData struct {
	InputText string
	ArtResult string
	Banner    string
	Error     string
}

// asciiArtHandler processes the ASCII art request and displays the result page
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	inputText := r.FormValue("text")
	banner := r.FormValue("banner")

	artResult, err := GenerateAsciiArt(inputText, banner)
	data := ArtData{InputText: inputText, Banner: banner, Error: ""}
	if err != nil {
		data.Error = err.Error()
	} else {
		data.ArtResult = artResult
	}

	// Render the result page with the generated ASCII art
	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
