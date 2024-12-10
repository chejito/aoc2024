package days

import (
	"aoc2024/helpers"
	"fmt"
	"strconv"
)

func RunDay9(lines []string) {
	fmt.Println("..: Day 9 solutions :..")
	line := lines[0]
	fmt.Println(line)
	day9Part1(line)
	//day9Part2()
}

func day9Part1(line string) {
	fmt.Println(".: Part 1 :.")
	parsedLine := getParsedLine(line)
	fmt.Printf("Parsed disk map:\n %s\n", parsedLine)
	compressedLine := compressLine(parsedLine)
	fmt.Printf("Compressed disk map:\n %s\n", compressedLine)
	var checksum int = getChecksum(compressedLine)
	fmt.Printf("Disk checksum: %d\n", checksum)
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

func getChecksum(lineArray []string) int {
	sum := 0
	for i := 0; i < len(lineArray); i++ {
		value := helpers.StringToInt(lineArray[i])
		sum += value * i
	}
	return sum
}

//func day9Part2() {
//	fmt.Println(".: Part 2 :.")
//}
