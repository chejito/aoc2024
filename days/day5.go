package days

import (
	"aoc2024/helpers"
	"fmt"
	"strings"
)

var processedOrders = make(map[string][]string)
var processedUpdates, orderedResults, unorderedResults, reorderedResults [][]string

func RunDay5(lines []string) {
	fmt.Println("..: Day 5 solutions :..")
	separatorIndex := helpers.GetIndex(lines, "")
	orders := lines[:separatorIndex]
	updates := lines[separatorIndex+1:]
	processOrders(orders)
	processUpdates(updates)

	fmt.Println("--- Orders ---")
	for key, order := range processedOrders {
		fmt.Printf("%s - %s\n", key, order)
	}

	fmt.Println("--- Updates ---")
	for _, update := range processedUpdates {
		fmt.Println(update)
	}
	orderedResults, unorderedResults = filterUpdates(processedUpdates)

	day5Part1()
	day5Part2()
}

func day5Part1() {
	fmt.Println(".: Part 1 :.")

	fmt.Println("--- Ordered Updates ---")
	for _, update := range orderedResults {
		fmt.Println(update)
	}

	sum := 0
	for _, result := range orderedResults {
		middlePage := result[len(result)/2]
		sum += helpers.StringToInt(middlePage)
	}
	fmt.Printf("Middle page values sum: %d\n", sum)
}

func day5Part2() {
	fmt.Println(".: Part 2 :.")
	fmt.Println("--- Unordered Updates ---")
	for _, update := range unorderedResults {
		fmt.Println(update)
	}
	reorderUpdates(unorderedResults)

	fmt.Println("--- Reordered Updates ---")
	for _, update := range reorderedResults {
		fmt.Println(update)
	}

	sum := 0
	for _, result := range reorderedResults {
		middlePage := result[len(result)/2]
		sum += helpers.StringToInt(middlePage)
	}
	fmt.Printf("Middle page values sum: %d\n", sum)
}

func processOrders(orders []string) {
	for _, order := range orders {
		orderPair := strings.Split(order, "|")
		processedOrders[orderPair[0]] = append(processedOrders[orderPair[0]], orderPair[1])
	}
}

func processUpdates(updates []string) {
	for _, update := range updates {
		item := strings.Split(update, ",")
		processedUpdates = append(processedUpdates, item)
	}
}

func filterUpdates(updates [][]string) ([][]string, [][]string) {
	var ordered, unordered [][]string
	for _, update := range updates {
		isOrdered := isUpdateOrdered(update)
		if isOrdered {
			ordered = append(ordered, update)
		} else {
			unordered = append(unordered, update)
		}
	}
	return ordered, unordered
}

func isUpdateOrdered(update []string) bool {
	for i, item := range update {
		firstItemOrders := processedOrders[item]
		left := update[i+1:]
		if len(left) > 0 && !containsAll(firstItemOrders, left) {
			return false
		}
	}
	return true
}

func containsAll(orders, items []string) bool {
	orderMap := make(map[string]bool)
	for _, order := range orders {
		orderMap[order] = true
	}

	for _, item := range items {
		if !orderMap[item] {
			return false
		}
	}
	return true
}

func reorderUpdates(results [][]string) {
	for _, result := range results {
		orderedResult := reorderUpdate(result)
		for !isUpdateOrdered(orderedResult) {
			orderedResult = reorderUpdate(orderedResult)
		}
		reorderedResults = append(reorderedResults, orderedResult)
	}
}

func reorderUpdate(update []string) []string {
	var result []string
	for i, item := range update {
		firstItemOrders := processedOrders[item]
		left := update[i+1:]
		if len(left) > 0 && !containsAll(firstItemOrders, left) {
			result = append(result, update[i+1])
			update[i+1] = item
		} else {
			result = append(result, item)
		}
	}
	return result
}
