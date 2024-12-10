package days

import (
	"aoc2024/helpers"
	"fmt"
	"strconv"
)

func RunDay9(lines []string) {
	fmt.Println("=== Day 9 solutions ===")
	line := lines[0]
	//fmt.Println(line)
	day9Part1(line)
	day9Part2(line)
}

func day9Part1(line string) {
	fmt.Println(".: Part 1 :.")
	parsedLine := getParsedLine(line)
	//fmt.Printf("Parsed disk map:\n %s\n", parsedLine)
	compressedLine := compressLine(parsedLine)
	//fmt.Printf("Compressed disk map:\n %s\n", compressedLine)
	checksum := getChecksum(compressedLine)
	fmt.Printf("Solution: %d\n", checksum)
}

func day9Part2(line string) {
	fmt.Println(".: Part 2 :.")
	parsedArray := getParsedArray(line)
	//fmt.Printf("Parsed disk map:\n %s\n", parsedArray)
	compressedArray := compressArray(parsedArray)
	//fmt.Printf("Compressed disk map:\n %s\n", compressedArray)
	filteredArray := filterSpaces(compressedArray)
	//fmt.Printf("Filtered disk map:\n %s\n", filteredArray)
	checksum := getChecksum(filteredArray)
	fmt.Printf("Solution: %d\n", checksum)
}

func getParsedLine(line string) []string {
	var parsedLine []string
	lineArray := helpers.StringToArrayNoSep(line)
	id := 0
	for i := 0; i < len(lineArray); i++ {
		times := helpers.StringToInt(lineArray[i])
		if i%2 == 0 {
			sequence := getSequenceData(times, id)
			parsedLine = append(parsedLine, sequence...)
			id++
		} else {
			sequence := getSequenceSpace(times)
			parsedLine = append(parsedLine, sequence...)
		}
	}
	return parsedLine
}

func getParsedArray(line string) [][]string {
	var parsedArray [][]string
	lineArray := helpers.StringToArrayNoSep(line)
	id := 0
	for i := 0; i < len(lineArray); i++ {
		times := helpers.StringToInt(lineArray[i])
		if i%2 == 0 {
			sequence := getSequenceData(times, id)
			parsedArray = append(parsedArray, sequence)
			id++
		} else {
			sequence := getSequenceSpace(times)
			if len(sequence) > 0 {
				parsedArray = append(parsedArray, sequence)
			}
		}
	}
	return parsedArray
}

func getSequenceData(times int, id int) []string {
	var sequence []string
	for i := 0; i < times; i++ {
		sequence = append(sequence, strconv.Itoa(id))
	}
	return sequence
}

func getSequenceSpace(times int) []string {
	var sequence []string
	for i := 0; i < times; i++ {
		sequence = append(sequence, ".")
	}
	return sequence
}

func compressLine(lineArray []string) []string {
	var compressedLine []string
	currentPosition := 0
	lastPosition := len(lineArray) - 1

	for i := currentPosition; i < lastPosition; i++ {
		if lineArray[i] != "." {
			compressedLine = append(compressedLine, lineArray[i])

		} else {
			if lineArray[lastPosition] != "." {
				compressedLine = append(compressedLine, lineArray[lastPosition])
				lineArray[lastPosition] = "."
				lastPosition--
			} else {
				for j := lastPosition; j >= i; j-- {
					if lineArray[j] != "." {
						compressedLine = append(compressedLine, lineArray[j])
						lineArray[j] = "."
						lastPosition--
						break
					}
				}
			}
		}
	}

	return compressedLine
}

func compressArray(lineArray [][]string) [][]string {
	lastPosition := len(lineArray) - 1

	for i := lastPosition; i >= 0; i-- {
		if isDigitBlock(lineArray[i]) {
			for j := 0; j < i; j++ {
				if !isDigitBlock(lineArray[j]) {
					isFree, index := hasFreeSpace(lineArray[j], len(lineArray[i]))
					if isFree {
						for k := index; k < index+len(lineArray[i]); k++ {
							lineArray[j][k] = lineArray[i][0]
						}
						for l := 0; l < len(lineArray[i]); l++ {
							lineArray[i][l] = "."
						}
						break
					}
				}
			}
		}
	}

	return lineArray
}

func hasFreeSpace(block []string, spaceNeeded int) (bool, int) {
	result := 0
	firstIndex := -1
	for i := 0; i < len(block); i++ {
		if block[i] == "." {
			if firstIndex == -1 {
				firstIndex = i
			}
			result++
		}
	}
	return result >= spaceNeeded, firstIndex
}

func isDigitBlock(block []string) bool {
	result := true
	for i := 0; i < len(block); i++ {
		if block[i] == "." {
			result = false
			break
		}
	}
	return result
}

func getChecksum(lineArray []string) int {
	sum := 0
	for i := 0; i < len(lineArray); i++ {
		if lineArray[i] != "." {
			value := helpers.StringToInt(lineArray[i])
			sum += value * i
		}
	}
	return sum
}

func filterSpaces(array [][]string) []string {
	var filteredArray []string
	for i := 0; i < len(array); i++ {
		filteredArray = append(filteredArray, array[i]...)
	}
	return filteredArray
}
