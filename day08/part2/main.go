package main

import (
	"crossbow.de/aoc2023/util"
	_ "embed"
	"fmt"
	"slices"
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

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func findStepsForEndingsWithZ(network map[string][]string, startingNodeAddress string, instructions []string) []int {
	currentNodeAddress := startingNodeAddress
	neededSteps := make([]int, 0)
	for steps := 0; true; steps++ {
		if steps > 30000 {
			fmt.Println("breaking on loops count")
			break
		}
		directionIndex := steps % len(instructions)
		directionAccessor := 0
		if instructions[directionIndex] == "R" {
			directionAccessor = 1
		}
		//fmt.Println(currentNodeAddress, steps, instructions[directionIndex])
		currentNodeAddress = network[currentNodeAddress][directionAccessor]
		if strings.HasSuffix(currentNodeAddress, "Z") {
			neededSteps = append(neededSteps, steps+1)
			//fmt.Println(currentNodeAddress, steps+1)
		}
	}
	return neededSteps
}

func process(input string) int {
	network := make(map[string][]string)
	instructions := make([]string, 0)
	startingNodeAddresses := make([]string, 0)

	lines := strings.Split(input, "\n")
	for lineIndex, line := range lines {
		if lineIndex == 0 {
			// instructions
			instructions = strings.Split(line, "")
			//fmt.Println(instructions)
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

		if strings.HasSuffix(address, "A") {
			startingNodeAddresses = append(startingNodeAddresses, address)
		}
	}
	fmt.Println("starting to iterate through ", startingNodeAddresses)

	minimumSteps := make([]int, 0)
	for _, startingNodeAddress := range startingNodeAddresses {
		foundSteps := findStepsForEndingsWithZ(network, startingNodeAddress, instructions)
		minimumSteps = append(minimumSteps, slices.Min(foundSteps))
	}
	fmt.Println("new steps with ", minimumSteps)
	// find lcm
	lcm := minimumSteps[0]
	for _, step := range minimumSteps[1:] {
		lcm = lcm * step / GCD(lcm, step)
	}

	return lcm
}

func main() {
	fmt.Println("Starting day08/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
