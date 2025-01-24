package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	inputFile := "UNSC-sample.txt"
	outputDir := "resolutions"

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	// Regular expression to match the pattern S/RES/#### (####)
	re := regexp.MustCompile(`S/RES/(\d{4}) \((\d{4})\)`)

	scanner := bufio.NewScanner(file)
	page := 1
	var currentText string
	var currentFilename string

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			// If there's already text collected, write it to the previous file
			if currentText != "" && currentFilename != "" {
				err := os.WriteFile(currentFilename, []byte(currentText), 0644)
				if err != nil {
					fmt.Println("Error writing file:", err)
					return
				}
				currentText = ""
			}

			// Prepare the new filename
			resolutionNumber := matches[1]
			year := matches[2]
			pageStr := fmt.Sprintf("%04d", page)
			currentFilename = filepath.Join(outputDir, fmt.Sprintf("%s-S-RES-%s-%s.txt", pageStr, resolutionNumber, year))
			page++
		}

		// Collect the text
		currentText += line + "\n"
	}

	// Write the last collected text to the last file
	if currentText != "" && currentFilename != "" {
		err := os.WriteFile(currentFilename, []byte(currentText), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
