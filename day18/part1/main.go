package main

import (
	"crossbow.de/aoc2023/util"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func process(input string) int {
	grid := strings.Split(input, "\n")
	return len(grid)
}

func main() {
	fmt.Println("Starting day16/part1")
	inputText := util.ReadFile("input.txt")
	output := process(inputText)
	fmt.Println(output)
	return
}
