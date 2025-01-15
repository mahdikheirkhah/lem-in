package fileHandler

import (
	"LemIn/errorHandler"
	"bufio"
	"os"
)

func ReadAll(fileName string) []string {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		errorHandler.CheckError(err, true)
		return nil
	}

	defer file.Close()

	// Read lines using a scanner
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		errorHandler.CheckError(err, true)
		return nil
	}

	return lines
}
