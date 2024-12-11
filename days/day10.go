package days

import (
	"aoc2024/helpers"
	"fmt"
)

var trailDirections = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func RunDay10(lines []string) {
	fmt.Println("=== Day 9 solutions ===")
	grid := helpers.ArrayOfStringNoSepToArrayOfArrayOfInt(lines)

	day10Part1(grid)
	day10Part2(grid)
}

func day10Part1(grid [][]int) {
	fmt.Println(".: Part 1 :.")
	trailheads := findTrailheads(grid)
	sum := 0
	for _, count := range trailheads {
		sum += count
	}
	fmt.Printf("Solution: %d\n", sum)
}

func day10Part2(grid [][]int) {
	fmt.Println(".: Part 2 :.")
	trailheads := findTrailheadsEveryTrails(grid)
	sum := 0
	for _, trailhead := range trailheads {
		sum += len(trailhead)
	}
	fmt.Printf("Solution: %d\n", sum)
}

func findTrailheads(grid [][]int) map[[2]int]int {
	rows := len(grid)
	cols := len(grid[0])
	result := make(map[[2]int]int) // Stores the amount of 9 cells hit by every 0 cell

	// BFS for every 0 cell
	bfs := func(startX, startY int) int {
		visited := make(map[[2]int]bool)
		queue := [][3]int{{startX, startY, 0}} // [x, y, currentValue]
		uniqueTargets := make(map[[2]int]bool)

		for len(queue) > 0 {
			x, y, currentValue := queue[0][0], queue[0][1], queue[0][2]
			queue = queue[1:]

			// If we reach a 9, register as hit
			if grid[x][y] == 9 {
				uniqueTargets[[2]int{x, y}] = true
				continue
			}

			// Mark as visited
			visited[[2]int{x, y}] = true

			// Explore valid neighbours
			for _, dir := range trailDirections {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
					if grid[nx][ny] == currentValue+1 && !visited[[2]int{nx, ny}] {
						queue = append(queue, [3]int{nx, ny, currentValue + 1})
						visited[[2]int{nx, ny}] = true
					}
				}
			}
		}

		return len(uniqueTargets)
	}

	// Start BFS from every cell that contains a 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				result[[2]int{i, j}] = bfs(i, j)
			}
		}
	}

	return result
}

func findTrailheadsEveryTrails(grid [][]int) map[[2]int][][][2]int {
	rows := len(grid)
	cols := len(grid[0])
	result := make(map[[2]int][][][2]int) //Stores every unique path from 0 cells to 9 cell

	// BFS for every 0 cell
	bfs := func(startX, startY int) [][][2]int {
		queue := [][3]interface{}{{startX, startY, [][2]int{{startX, startY}}}} // [x, y, path]
		var paths [][][2]int

		for len(queue) > 0 {
			x, y, path := queue[0][0].(int), queue[0][1].(int), queue[0][2].([][2]int)
			queue = queue[1:]

			// If we reach a 9, add path to the result list
			if grid[x][y] == 9 {
				paths = append(paths, append([][2]int{}, path...))
				continue
			}

			// Explore valid neighbours
			for _, dir := range trailDirections {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
					if grid[nx][ny] == grid[x][y]+1 {
						newPath := append([][2]int{}, path...)
						newPath = append(newPath, [2]int{nx, ny})
						queue = append(queue, [3]interface{}{nx, ny, newPath})
					}
				}
			}
		}

		return paths
	}

	// Start BFS from every cell that contains a 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				result[[2]int{i, j}] = bfs(i, j)
			}
		}
	}

	return result
}
