package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
