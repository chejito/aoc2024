package day1

import (
	"aoc2024/helpers"
	"fmt"
	"sort"
)

var leftList []int
var rightList []int
var distanceList1 []int
var distanceList2 []int

func Run(lines []string) {
	fmt.Println("..: Day 1 solutions :..")

	for _, line := range lines {
		array := helpers.StringToArray(line)
		leftList = append(leftList, helpers.GetIntFromString(array[0]))
		rightList = append(rightList, helpers.GetIntFromString(array[1]))
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	part1()
	part2()
}

func part1() {
	fmt.Println(".: Part 1 :.")

	for i := 0; i < len(leftList); i++ {
		distance := getDistancePart1(leftList[i], rightList[i])
		distanceList1 = append(distanceList1, distance)
	}

	printResult(distanceList1)
}

func part2() {
	fmt.Println(".: Part 2 :.")

	for i := 0; i < len(leftList); i++ {
		distance := getDistancePart2(leftList[i], rightList)
		distanceList2 = append(distanceList2, distance)
	}

	printResult(distanceList2)
}

func getDistancePart1(leftValue int, rightValue int) int {
	if leftValue > rightValue {
		return leftValue - rightValue
	}
	return rightValue - leftValue
}

func getDistancePart2(leftValue int, rightList []int) int {
	var timesInRightList int = getTimesInRightList(leftValue, rightList)
	return leftValue * timesInRightList
}

func getTimesInRightList(leftValue int, rightList []int) int {
	times := 0
	for _, rightValue := range rightList {
		if rightValue == leftValue {
			times++
		}
	}
	return times
}

func printResult(distanceList []int) {
	totalDistance := 0
	for _, num := range distanceList {
		totalDistance += num
	}
	fmt.Printf("Total distance is %d\n", totalDistance)
}
