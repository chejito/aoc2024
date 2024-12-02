package main

import (
	"aoc2024/days"
	"aoc2024/helpers"
	"fmt"
	"os"
)

var lines []string
var dayFunctions = map[string]func(){
	"1":  runDay1,
	"t1": runDay1,
	"2":  runDay2,
	"t2": runDay2,
}

func runDay1() {
	days.RunDay1(lines)
}

func runDay2() {
	days.RunDay2(lines)
}

func main() {
	config := helpers.ReadConfig()

	if len(os.Args) < 2 {
		fmt.Println("Please, enter day as a parameter.")
		return
	}

	day := os.Args[1]
	path := config[day]

	lines = helpers.ReadDayFile(path)

	if method, exists := dayFunctions[day]; exists {
		method()
	} else {
		fmt.Printf("Did not find day '%s'\n", day)
	}
}
