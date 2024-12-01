package main

import (
	"aoc2024/day1"
	"aoc2024/helpers"
	"fmt"
	"os"
)

var lines []string
var dayFunctions = map[string]func(){
	"1":  runDay1,
	"t1": runDay1,
}

func runDay1() {
	day1.Run(lines)
}

func main() {
	days := helpers.ReadConfig()

	if len(os.Args) < 2 {
		fmt.Println("Please, enter day as a parameter.")
		return
	}

	day := os.Args[1]
	path := days[day]

	lines = helpers.ReadDayFile(path)

	if method, exists := dayFunctions[day]; exists {
		method()
	} else {
		fmt.Printf("Did not find day '%s'\n", day)
	}
}
