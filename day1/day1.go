package day1

import (
	"aoc2024/helpers"
	"fmt"
	"sort"
)

var list1 []int
var list2 []int
var distanceList1 []int
var distanceList2 []int

func Run(lines []string) {

	for _, line := range lines {
		array := helpers.StringToArray(line)
		list1 = append(list1, helpers.GetIntFromString(array[0]))
		list2 = append(list2, helpers.GetIntFromString(array[1]))
	}

	sort.Ints(list1)
	sort.Ints(list2)

	part1()
	part2()
}

func part1() {
	fmt.Println(".: Part 1 :.")

	for i := 0; i < len(list1); i++ {
		distance := getDistancePart1(list1[i], list2[i])
		distanceList1 = append(distanceList1, distance)
	}

	printResult(distanceList1)
}

func part2() {
	fmt.Println(".: Part 2 :.")

	for i := 0; i < len(list1); i++ {
		distance := getDistancePart2(list1[i], list2)
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
