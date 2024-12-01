package utils

import (
	"bufio"
	"fmt"
	"os"
)

func LoadFile() []string {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open the file: %v", err)
	}
	defer file.Close()

	// Create a slice to hold lines
	var lines []string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Append each line to the slice
		lines = append(lines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading the file: %v", err)
	}

	// Print the resulting list
	return lines
}
