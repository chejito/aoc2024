package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func ReadDayFile(path string) []string {

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil
	}

	return lines
}
