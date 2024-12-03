package days

import (
	"aoc2024/helpers"
	"fmt"
	"math"
)

func RunDay2(lines []string) {
	fmt.Println("..: Day 2 solutions :..")

	levels := helpers.ArrayOfStringToArrayOfArrayOfInt(lines)

	day2Part1(levels)
	day2Part2(levels)
}

func day2Part1(levels [][]int) {
	fmt.Println(".: Part 1 :.")

	safeReports := 0
	for _, level := range levels {
		if isSafeReport(level) {
			safeReports++
		}
	}
	fmt.Printf("Total safe reports: %d\n", safeReports)
}

func day2Part2(levels [][]int) {
	fmt.Println(".: Part 2 :.")
	safeReports := 0
	for _, level := range levels {
		if isSafeReport(level) {
			safeReports++
		} else {
			if isToleratedSafety(level) {
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
