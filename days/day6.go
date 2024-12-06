package days

import (
	"aoc2024/helpers"
	"fmt"
)

type position struct {
	x int
	y int
}

var directions = []position{
	{x: -1, y: 0},
	{x: 0, y: 1},
	{x: 1, y: 0},
	{x: 0, y: -1},
}
var rows, cols int
var guardPosition position
var guardDirection int
var visitedPositions map[position]bool

func RunDay6(lines []string) {
	fmt.Println("..: Day 5 solutions :..")
	grid := helpers.ArrayOfStringToArrayOfArrayOfString(lines)
	rows, cols = len(grid), len(grid[0])
	fmt.Printf("Grid size: %d x %d\n", rows, cols)

	day6Part1(grid)
	//day6Part2()
}

func setInitialValues(grid [][]string) {
	guardDirection = 0
	visitedPositions = make(map[position]bool)

	for row := range grid {
		for col := range grid[row] {

			if grid[row][col] == "^" || grid[row][col] == "v" || grid[row][col] == "<" || grid[row][col] == ">" {
				guardPosition = position{x: row, y: col}
				visitedPositions[guardPosition] = true
				break
			}
		}
	}
}

func day6Part1(grid [][]string) {
	fmt.Println(".: Part 1 :.")
	setInitialValues(grid)
	fmt.Printf("Initial guard position at row %d, col %d\n", guardPosition.x, guardPosition.y)
	fmt.Printf("Initial visited positions: %d\n", len(visitedPositions))
	patrol(grid)
	fmt.Printf("Total visited positions: %d\n", len(visitedPositions))
}

func day6Part2() {
	fmt.Println(".: Part 2 :.")
}

func patrol(grid [][]string) {
	for {
		direction := directions[guardDirection]
		dx, dy := direction.x, direction.y
		nx, ny := guardPosition.x+dx, guardPosition.y+dy

		if !(nx >= 0 && nx < rows && ny >= 0 && ny < cols) {
			break
		}
		if grid[nx][ny] == "#" {
			guardDirection = (guardDirection + 1) % 4
		} else {
			guardPosition.x = nx
			guardPosition.y = ny
			visitedPositions[guardPosition] = true
		}
	}
}
