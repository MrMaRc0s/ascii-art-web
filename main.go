package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainPage)                 // Main page handler
	http.HandleFunc("/ascii-art", asciiArtHandler) // ASCII art generation handler
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
