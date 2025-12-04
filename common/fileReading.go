package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readGridFromFile reads a text file into a 2D slice of runes (Unicode characters).
func ReadGridFromFile(filePath string) ([][]rune, error) {
	// Open the file. It's important to close it using defer.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	// Read through the file line by line.
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		// Convert the string line into a slice of runes.
		// Using []rune handles multi-byte Unicode characters correctly.
		chars := []rune(line)
		grid = append(grid, chars)
	}

	// Check for errors during scanning. EOF is not reported as an error.
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error during file scanning: %w", err)
	}

	return grid, nil
}
