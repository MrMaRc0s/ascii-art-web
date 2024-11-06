package main

import (
	"strings"
)

func GenerateAsciiArt(inputText, banner string) (string, error) {
	lines, err := loadBanner(banner)
	if err != nil {
		return "", err
	}

	sepText := strings.Split(inputText, "\\n") // Handling literal "\n" from the input
	var result strings.Builder
	printAsciiArtRecursive(sepText, lines, &result)
	return result.String(), nil
}

func printAsciiArtRecursive(sentences []string, textFile []string, result *strings.Builder) {
	if len(sentences) == 0 {
		return
	}

	if sentences[0] != "" {
		printSentenceAsciiRecursive(sentences[0], textFile, 1, result)
		result.WriteString("\n")
	} else {
		result.WriteString("\n")
	}

	printAsciiArtRecursive(sentences[1:], textFile, result)
}

func printSentenceAsciiRecursive(word string, textFile []string, h int, result *strings.Builder) {
	if h > 8 {
		return
	}

	for i := 0; i < len(word); i++ {
		for lineIndex, line := range textFile {
			if lineIndex == (int(word[i])-32)*9+h {
				result.WriteString(line)
			}
		}
	}
	result.WriteString("\n")

	printSentenceAsciiRecursive(word, textFile, h+1, result)
}
