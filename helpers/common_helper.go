package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToArray(str string) []string {
	return strings.Fields(str)
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error when converting:", err)
	}
	return num
}

func ArrayOfStringToArrayOfArrayOfInt(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		array := StringToArray(line)
		intArray := StringArrayToIntArray(array)
		result = append(result, intArray)
	}
	return result
}

func ArrayOfStringNoSepToArrayOfArrayOfInt(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		array := StringToArrayNoSep(line)
		intArray := StringArrayToIntArray(array)
		result = append(result, intArray)
	}
	return result
}

func ArrayOfStringToArrayOfArrayOfString(lines []string) [][]string {
	var result [][]string
	for _, line := range lines {
		array := StringToArrayNoSep(line)
		result = append(result, array)
	}
	return result
}

func StringArrayToIntArray(array []string) []int {
	var result []int
	for _, value := range array {
		result = append(result, StringToInt(value))
	}
	return result
}

func GetIndex(array []string, itemToFind string) int {
	index := -1
	for i, item := range array {
		if item == itemToFind {
			index = i
		}
	}
	return index
}

func StringToArrayNoSep(line string) []string {
	return strings.Split(line, "")
}
