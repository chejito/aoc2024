package days

import (
	"aoc2024/helpers"
	"fmt"
	"math"
)

func RunDay2(lines []string) {
	fmt.Println("..: Day 2 solutions :..")

	day2Part1(lines)
	day2Part2(lines)
}

func day2Part1(lines []string) {
	fmt.Println(".: Part 1 :.")

	safeReports := 0
	for _, line := range lines {
		array := helpers.StringToArray(line)
		intArray := helpers.StringArrayToIntArray(array)
		if isSafeReport(intArray) {
			safeReports++
		}
	}
	fmt.Printf("Total safe reports: %d\n", safeReports)
}

func day2Part2(lines []string) {
	fmt.Println(".: Part 2 :.")
	safeReports := 0
	for _, line := range lines {
		array := helpers.StringToArray(line)
		intArray := helpers.StringArrayToIntArray(array)
		if isSafeReport(intArray) {
			safeReports++
		} else {
			if isToleratedSafety(intArray) {
				safeReports++
			}
		}
	}
	fmt.Printf("Total safe reports: %d\n", safeReports)
}

func isSafeReport(intArray []int) bool {
	allIncreasing := isAllIncreasing(intArray)
	if allIncreasing && isGradually(intArray) {
		return true
	}
	allDecreasing := isAllDecreasing(intArray)
	if allDecreasing && isGradually(intArray) {
		return true
	}
	return false
}

func isToleratedSafety(intArray []int) bool {
	for i := 0; i < len(intArray); i++ {
		sliceWithoutCurrent := make([]int, 0, len(intArray)-1)
		sliceWithoutCurrent = append(sliceWithoutCurrent, intArray[:i]...)
		sliceWithoutCurrent = append(sliceWithoutCurrent, intArray[i+1:]...)
		if isSafeReport(sliceWithoutCurrent) {
			return true
		}
	}
	return false
}

func isAllIncreasing(intArray []int) bool {
	for i := 1; i < len(intArray); i++ {
		if intArray[i] <= intArray[i-1] {
			return false
		}
	}
	return true
}

func isAllDecreasing(intArray []int) bool {
	for i := 1; i < len(intArray); i++ {
		if intArray[i] >= intArray[i-1] {
			return false
		}
	}
	return true
}

func isGradually(array []int) bool {
	for i := 1; i < len(array); i++ {
		value := math.Abs(float64(array[i] - array[i-1]))
		if value < 1 || value > 3 {
			return false
		}
	}
	return true
}
