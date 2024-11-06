package main

import (
	"net/http"
	"text/template"
)

// mainPage serves the main HTML page where users can input text and select banners
func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
