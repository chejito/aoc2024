package days

import (
	"aoc2024/helpers"
	"fmt"
)

func RunDay8(lines []string) {
	fmt.Println("..: Day 8 solutions :..")
	grid := helpers.ArrayOfStringToArrayOfArrayOfString(lines)
	rows, cols = len(grid), len(grid[0])
	antennaLocations := getAntennaLocations(grid)
	for _, location := range antennaLocations {
		fmt.Println(location)
	}
	day8Part1(antennaLocations)
	day8Part2(antennaLocations)
}

func getAntennaLocations(grid [][]string) map[string][][2]int {
	locations := make(map[string][][2]int)
	for row, line := range grid {
		for col, val := range line {
			if grid[row][col] != "." {
				locations[val] = append(locations[val], [2]int{row, col})
			}
		}
	}
	return locations
}

func day8Part1(antennaLocations map[string][][2]int) {
	fmt.Println(".: Part 1 :.")
	var antiNodeLocations = getAntiNodes("1", antennaLocations)
	fmt.Printf("Antinode locations: %v\n", antiNodeLocations)
	fmt.Printf("Total antinodes: %d\n", len(antiNodeLocations))
}

func day8Part2(antennaLocations map[string][][2]int) {
	fmt.Println(".: Part 2 :.")
	var antiNodeLocations = getAntiNodes("2", antennaLocations)
	fmt.Printf("Antinode locations: %v\n", antiNodeLocations)
	fmt.Printf("Total antinodes: %d\n", len(antiNodeLocations))
}

func getAntiNodes(part string, antennaLocations map[string][][2]int) [][2]int {
	nodesMap := make(map[[2]int]bool)
	antiNodes := make([][2]int, 0)
	for _, antennaTypeLocation := range antennaLocations {
		if len(antennaTypeLocation) > 1 {
			var unfilteredNodes [][2]int
			if part == "1" {
				unfilteredNodes = getAntiNodesForTypePart1(antennaTypeLocation)
			}
			if part == "2" {
				unfilteredNodes = getAntiNodesForTypePart2(antennaTypeLocation)
			}

			for _, node := range unfilteredNodes {
				if !nodesMap[node] {
					nodesMap[node] = true
					antiNodes = append(antiNodes, node)
				}
			}
		}
	}
	return antiNodes
}

func getAntiNodesForTypePart1(typeLocations [][2]int) [][2]int {
	nodes := make([][2]int, 0)
	var dfs func(first [2]int, others [][2]int)
	dfs = func(first [2]int, others [][2]int) {
		if len(others) > 0 {
			for _, location := range others {
				dx, dy := location[0]-first[0], location[1]-first[1]
				p1x, p1y := first[0]-dx, first[1]-dy
				if p1x >= 0 && p1x < rows && p1y >= 0 && p1y < cols {
					nodes = append(nodes, [2]int{p1x, p1y})
				}
				p2x, p2y := location[0]+dx, location[1]+dy
				if p2x >= 0 && p2x < rows && p2y >= 0 && p2y < cols {
					nodes = append(nodes, [2]int{p2x, p2y})
				}
			}
		}
		nextFirst := others[0]
		nextOthers := others[1:]
		if len(nextOthers) > 0 {
			dfs(nextFirst, nextOthers)
		}
	}
	first := typeLocations[0]
	others := typeLocations[1:]
	dfs(first, others)
	return nodes
}

func getAntiNodesForTypePart2(typeLocations [][2]int) [][2]int {
	nodes := make([][2]int, 0)
	var dfs func(first [2]int, others [][2]int)
	dfs = func(first [2]int, others [][2]int) {
		if len(others) > 0 {
			for _, location := range others {
				dx, dy := location[0]-first[0], location[1]-first[1]
				p1x, p1y := first[0]+dx, first[1]+dy
				for {
					nx, ny := p1x, p1y
					if !(nx >= 0 && nx < rows && ny >= 0 && ny < cols) {
						break
					}
					nodes = append(nodes, [2]int{nx, ny})
					p1x, p1y = nx+dx, ny+dy
				}

				p2x, p2y := location[0]-dx, location[1]-dy
				for {
					nx, ny := p2x, p2y
					if !(nx >= 0 && nx < rows && ny >= 0 && ny < cols) {
						break
					}
					nodes = append(nodes, [2]int{nx, ny})
					p2x, p2y = nx-dx, ny-dy
				}
			}
		}
		nextFirst := others[0]
		nextOthers := others[1:]
		if len(nextOthers) > 0 {
			dfs(nextFirst, nextOthers)
		}
	}
	first := typeLocations[0]
	others := typeLocations[1:]
	dfs(first, others)
	return nodes
}
