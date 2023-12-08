package main

import (
	"crossbow.de/aoc2023/util"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func process(input string) int {
	network := make(map[string][]string)
	instructions := make([]string, 0)
	startingNodeAddress := "AAA"
	endingNodeAddress := "ZZZ"

	lines := strings.Split(input, "\n")
	for lineIndex, line := range lines {
		if lineIndex == 0 {
			// instructions
			instructions = strings.Split(line, "")
			fmt.Println(instructions)
			continue
		}
		if lineIndex == 1 {
			continue
		}
		// handle node creation
		addressAndRoutes := strings.Split(line, " = ")
		address := addressAndRoutes[0]
		routes := strings.Split(strings.Trim(addressAndRoutes[1], "()"), ", ")
		network[address] = routes
	}

	currentNodeAddress := startingNodeAddress
	neededSteps := 0
	for steps := 0; true; steps++ {
		if steps > 100000000 {
			fmt.Println("breaking on loops count")
			break
		}
		directionIndex := steps % len(instructions)
		directionAccessor := 0
		if instructions[directionIndex] == "R" {
			directionAccessor = 1
		}
		fmt.Println(currentNodeAddress, steps, instructions[directionIndex])
		currentNodeAddress = network[currentNodeAddress][directionAccessor]
		if currentNodeAddress == endingNodeAddress {
			neededSteps = steps
			break
		}
	}

	return neededSteps + 1
}

func main() {
	fmt.Println("Starting day08/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
