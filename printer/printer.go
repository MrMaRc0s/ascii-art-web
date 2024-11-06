package printer

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GenerateAsciiArt generates ASCII art for the given text and banner
func GenerateAsciiArt(inputString, bannerName string) (string, error) {
	// Construct the file path based on the banner name
	bannerFile := fmt.Sprintf("banners/%s.txt", bannerName)

	// Read the banner file
	content, err := ioutil.ReadFile(bannerFile)
	if err != nil {
		return "", fmt.Errorf("failed to read banner file: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	// Create a map to hold the ASCII art for each character
	var blocks [][]string
	var currentBlock []string
	bannerMap := make(map[rune][]string)

	// Iterate through lines and group them into blocks
	for index, line := range lines {
		if index%9 == 0 {
			continue
		}
		currentBlock = append(currentBlock, line)
		if len(currentBlock) == 8 {
			blocks = append(blocks, currentBlock)
			currentBlock = []string{} // Reset for the next block
		}
	}

	// Map each block to its corresponding ASCII character (from 32 to 126)
	for i := 0; i < 95; i++ {
		bannerMap[rune(i+32)] = blocks[i]
	}

	var outputLines []string
	inputLines := strings.Split(inputString, "\\n")

	// Iterate over each input line and build ASCII art representation
	for _, line := range inputLines {
		if line == "" {
			outputLines = append(outputLines, "")
			continue
		}
		tempOutputLines := make([]string, 8)

		for _, char := range line {
			art, exists := bannerMap[char]
			if exists {
				for i := 0; i < 8; i++ {
					tempOutputLines[i] += art[i]
				}
			} else {
				fmt.Printf("Character '%c' not found in banner map\n", char)
			}
		}

		// Append the constructed temp lines to the final output
		outputLines = append(outputLines, tempOutputLines...)
	}

	return strings.Join(outputLines, "\n"), nil
}
