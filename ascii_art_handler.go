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
            // Return 404 Not Found if the request is not a POST
            http.NotFound(w, r)
            return
    	}

    // Parse form data
    err := r.ParseForm()
    	if err != nil {
            http.Error(w, "Bad Request - Unable to parse form data", http.StatusBadRequest)
            return
    }

    inputText := r.FormValue("text")
    banner := r.FormValue("banner")

    // Check if the provided banner is valid
    if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
        http.Error(w, "Bad Request - Invalid banner template", http.StatusBadRequest)
        return
    }

    artResult, err := GenerateAsciiArt(inputText, banner)
    data := ArtData{InputText: inputText, Banner: banner, Error: ""}
    if err != nil {
        if err.Error() == "banner file not found" {
            http.Error(w, "Not Found - Banner file missing", http.StatusNotFound)
        } else {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        }
        return
    }

    data.ArtResult = artResult
    tmpl, err := template.ParseFiles("templates/result.html")
    if err != nil {
        http.Error(w, "Internal Server Error - Template parsing failed", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK) // Explicitly set status code 200
    tmpl.Execute(w, data)
}
