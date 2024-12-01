package day1

import (
	"aoc2024/helpers"
	"fmt"
	"sort"
)

var list1 []int
var list2 []int
var distanceList []int

func Run(lines []string) {
	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
	fmt.Println(".: Part 1 :.")

	for _, line := range lines {
		array := helpers.StringToArray(line)
		list1 = append(list1, helpers.GetIntFromString(array[0]))
		list2 = append(list2, helpers.GetIntFromString(array[1]))
	}

	sort.Ints(list1)
	sort.Ints(list2)

	fmt.Println(list1)
	fmt.Println(list2)

	for i := 0; i < len(list1); i++ {
		distance := GetDistance(list1[i], list2[i])
		distanceList = append(distanceList, distance)
	}

	fmt.Println(distanceList)
	sum := 0
	for _, num := range distanceList {
		sum += num
	}
	fmt.Printf("Total distance is %d\n", sum)

	fmt.Println(".: Part 2 :.")
}

func GetDistance(i1 int, i2 int) int {
	if i1 > i2 {
		return i1 - i2
	}
	return i2 - i1
}
