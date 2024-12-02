package days

import (
	"aoc2024/helpers"
	"fmt"
	"math"
)

func RunDay2(lines []string) {
	fmt.Println("..: Day 2 solutions :..")
	for _, line := range lines {
		fmt.Println(line)
	}
	day2Part1(lines)
	day2Part2()
}

func day2Part1(lines []string) {
	fmt.Println(".: Part 1 :.")
	safeReports := 0
	for _, line := range lines {
		if isSafeReport(line) {
			safeReports++
		}
	}
	fmt.Printf("Total safe reports: %d\n", safeReports)
}

func day2Part2() {
	fmt.Println(".: Part 2 :.")
}

func isSafeReport(line string) bool {
	array := helpers.StringToArray(line)
	intArray := helpers.StringArrayToIntArray(array)
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

func isAllIncreasing(array []int) bool {
	for i := 1; i < len(array); i++ {
		if array[i] <= array[i-1] {
			return false
		}
	}
	return true
}

func isAllDecreasing(array []int) bool {
	for i := 1; i < len(array); i++ {
		if array[i] >= array[i-1] {
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
