package days

import (
	"aoc2024/helpers"
	"fmt"
	"regexp"
)

func RunDay3(lines []string) {
	program := ""
	for _, line := range lines {
		program += line
	}
	day3Part1(program)
	day3Part2(program)
}

func day3Part1(program string) {
	fmt.Println(".: Part 1 :.")
	product := processLine(program)
	fmt.Printf("Total mul instructions product: %d\n", product)
}

func day3Part2(program string) {
	fmt.Println(".: Part 2 :.")
	cleanProgram := removeDonts(program)
	product := processLine(cleanProgram)
	fmt.Printf("Total mul instructions product: %d\n", product)
}

func removeDonts(line string) string {
	pattern := `don't\(\).*?do\(\)|don't\(\).*`
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(line, "")
	return result
}

func processLine(line string) int {
	result := 0
	pattern := `(?s)mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		a := helpers.StringToInt(match[1])
		b := helpers.StringToInt(match[2])

		var product = a * b
		result += product
	}
	return result
}
