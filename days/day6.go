package days

import (
	"aoc2024/helpers"
	"fmt"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var directions2 = [4][2]int{
	{-1, 0}, // Norte
	{0, 1},  // Este
	{1, 0},  // Sur
	{0, -1}, // Oeste
}

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
	fmt.Println("..: Day 6 solutions :..")
	grid := helpers.ArrayOfStringToArrayOfArrayOfString(lines)
	rows, cols = len(grid), len(grid[0])
	fmt.Printf("Grid size: %d x %d\n", rows, cols)

	day6Part1(grid)
	day6Part2(grid)
}

func day6Part1(grid [][]string) {
	fmt.Println(".: Part 1 :.")
	setInitialValues(grid)
	fmt.Printf("Initial guard position at row %d, col %d\n", guardPosition.x, guardPosition.y)
	fmt.Printf("Initial visited positions: %d\n", len(visitedPositions))
	patrol(grid)
	fmt.Printf("Total visited positions: %d\n", len(visitedPositions))
}

func day6Part2(grid [][]string) {
	fmt.Println(".: Part 2 :.")
	setInitialValues(grid)
	totalLoopPositions := countLoopPositions(grid)
	fmt.Printf("Total loop positions: %d\n", totalLoopPositions)
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

func patrol(grid [][]string) {
	visitedPositions = make(map[position]bool)

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

func isInLoop(grid [][]string, startPos [2]int, startDir Direction) bool {
	visitedStates := make(map[[3]int]bool)

	x, y := startPos[0], startPos[1]
	directionIdx := int(startDir)

	for {
		state := [3]int{x, y, directionIdx}
		if visitedStates[state] {
			return true
		}
		visitedStates[state] = true

		dx, dy := directions2[directionIdx][0], directions2[directionIdx][1]
		nx, ny := x+dx, y+dy

		if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
			return false
		}
		if grid[nx][ny] == "#" {
			directionIdx = (directionIdx + 1) % 4
		} else {
			x, y = nx, ny
		}
	}
}

func countLoopPositions(grid [][]string) int {
	loopPositions := 0

	var startPos [2]int
	var startDir Direction

	for row := range rows {
		for col := range cols {
			switch grid[row][col] {
			case "^":
				startPos = [2]int{row, col}
				startDir = North
				grid[row][col] = "."
			case ">":
				startPos = [2]int{row, col}
				startDir = East
				grid[row][col] = "."
			case "v":
				startPos = [2]int{row, col}
				startDir = South
				grid[row][col] = "."
			case "<":
				startPos = [2]int{row, col}
				startDir = West
				grid[row][col] = "."
			}
		}
	}

	for row := range rows {
		for col := range cols {
			if grid[row][col] == "." {
				grid[row][col] = "#"
				if isInLoop(grid, startPos, startDir) {
					loopPositions++
				}
				grid[row][col] = "."
			}
		}
	}

	return loopPositions
}
