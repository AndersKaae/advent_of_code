package utils

import (
	"bufio"
	"fmt"
	"os"
)

func LoadFile(filename string) []string {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		err := fmt.Errorf("Failed to open the file: %v", err)
		panic(err)

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
