package days

import "fmt"

var xmas map[string]bool = map[string]bool{
	"XMAS": true,
	"SAMX": true,
}

var shapes map[string]bool = map[string]bool{
	"SSAMM": true,
	"MMASS": true,
	"MSAMS": true,
	"SMASM": true,
}

func RunDay4(lines []string) {
	fmt.Println("..: Day 4 solutions :..")
	day4Part1(lines)
	day4Part2(lines)
}

func day4Part1(lines []string) {
	fmt.Println(".: Part 1 :.")

	horizontals := getHorizontals(lines)
	verticals := getVerticals(lines)
	nwSeDiagonals := getNwSeDiagonals(lines)
	neSwDiagonals := getNeSwDiagonals(lines)
	sum := horizontals + verticals + nwSeDiagonals + neSwDiagonals

	fmt.Printf("Total XMAS words: %d\n", sum)
}

func day4Part2(lines []string) {
	fmt.Println(".: Part 2 :.")
	totalShapes := getXmasShape(lines)
	fmt.Printf("Total XMAS words: %d\n", totalShapes)
}

func getHorizontals(lines []string) int {
	sum := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines)-3; j++ {
			section := getHorizontalSection(lines, i, j)
			if xmas[section] {
				sum++
			}
		}
	}
	return sum
}

func getVerticals(lines []string) int {
	sum := 0
	for i := 0; i < len(lines)-3; i++ {
		for j := 0; j < len(lines); j++ {
			section := getVerticalSection(lines, i, j)
			if xmas[section] {
				sum++
			}
		}
	}
	return sum
}

func getNwSeDiagonals(lines []string) int {
	sum := 0
	for i := 0; i < len(lines)-3; i++ {
		for j := 0; j < len(lines)-3; j++ {
			section := getNwSeDiagonalSection(lines, i, j)
			if xmas[section] {
				sum++
			}
		}
	}
	return sum
}

func getNeSwDiagonals(lines []string) int {
	sum := 0
	for i := 0; i < len(lines)-3; i++ {
		for j := 3; j < len(lines); j++ {
			section := getNeSwDiagonalSection(lines, i, j)
			if xmas[section] {
				sum++
			}
		}
	}
	return sum
}

func getXmasShape(lines []string) int {
	sum := 0
	for i := 0; i < len(lines)-2; i++ {
		for j := 0; j < len(lines)-2; j++ {
			section := GetShapeSection(lines, i, j)
			if shapes[section] {
				sum++
			}
		}
	}
	return sum
}

func getHorizontalSection(lines []string, i int, j int) string {
	runes := []rune{rune(lines[i][j]), rune(lines[i][j+1]), rune(lines[i][j+2]), rune(lines[i][j+3])}
	return string(runes)
}

func getVerticalSection(lines []string, i int, j int) string {
	runes := []rune{rune(lines[i][j]), rune(lines[i+1][j]), rune(lines[i+2][j]), rune(lines[i+3][j])}
	return string(runes)
}

func getNwSeDiagonalSection(lines []string, i int, j int) string {
	runes := []rune{rune(lines[i][j]), rune(lines[i+1][j+1]), rune(lines[i+2][j+2]), rune(lines[i+3][j+3])}
	return string(runes)
}

func getNeSwDiagonalSection(lines []string, i int, j int) string {
	runes := []rune{rune(lines[i][j]), rune(lines[i+1][j-1]), rune(lines[i+2][j-2]), rune(lines[i+3][j-3])}
	return string(runes)
}

func GetShapeSection(lines []string, i int, j int) string {
	runes := []rune{rune(lines[i][j]), rune(lines[i][j+2]), rune(lines[i+1][j+1]), rune(lines[i+2][j]), rune(lines[i+2][j+2])}
	return string(runes)
}
