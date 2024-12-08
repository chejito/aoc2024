package days

import (
	"aoc2024/helpers"
	"fmt"
	"strings"
)

var equations [][]int

func RunDay7(lines []string) {
	fmt.Println("..: Day 7 solutions :..")

	for _, line := range lines {
		equation := strings.Replace(line, ":", "", -1)
		stringEquations := strings.Split(equation, " ")
		var intEquations []int
		for _, number := range stringEquations {
			intEquations = append(intEquations, helpers.StringToInt(number))
		}
		equations = append(equations, intEquations)
	}

	for _, equation := range equations {
		fmt.Printf("Value: %d - Numbers: %d\n", equation[0], equation[1:])
	}

	day7Part1()
	day7Part2()
}

func day7Part1() {
	fmt.Println(".: Part 1 :.")
	calibrationResult := getCalibrationResult("1")
	fmt.Printf("Calibration result: %d\n", calibrationResult)
}

func day7Part2() {
	fmt.Println(".: Part 2 :.")
	calibrationResult := getCalibrationResult("2")
	fmt.Printf("Calibration result: %d\n", calibrationResult)
}

func getCalibrationResult(part string) int {
	result := 0
	for _, equation := range equations {
		value := equation[0]
		numbers := equation[1:]
		if part == "1" {
			if isSolvable(value, numbers) {
				result += value
			}
		}
		if part == "2" {
			if isSolvableWithConcatenation(value, numbers) {
				result += value
			}
		}
	}
	return result
}

func isSolvable(value int, numbers []int) bool {
	if len(numbers) == 0 {
		return false
	}
	var dfs func(index, current int) bool
	dfs = func(index, current int) bool {
		if index == len(numbers) {
			return current == value
		}

		next := numbers[index]
		return dfs(index+1, current+next) || dfs(index+1, current*next)
	}

	return dfs(1, numbers[0])
}

func isSolvableWithConcatenation(value int, numbers []int) bool {
	if len(numbers) == 0 {
		return false
	}
	var dfs func(index, current int) bool
	dfs = func(index, current int) bool {
		if index == len(numbers) {
			return current == value
		}

		next := numbers[index]
		return dfs(index+1, current+next) || dfs(index+1, current*next) || dfs(index+1, concatenate(current, next))
	}

	return dfs(1, numbers[0])
}

func concatenate(current, next int) int {
	concatenation := fmt.Sprintf("%d%d", current, next)
	return helpers.StringToInt(concatenation)
}
