package main

import (
	"crossbow.de/aoc2023/util"
	_ "embed"
	"fmt"
	"strconv"
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

func processSequence(array []int) int {
	if allElementsAreZeros(array) {
		return 0
	}
	deltas := make([]int, len(array)-1)
	for i := 0; i < len(array)-1; i++ {
		deltas[i] = array[i+1] - array[i]
	}
	diff := processSequence(deltas)
	return array[len(array)-1] + diff
}

func allElementsAreZeros(array []int) bool {
	for _, x := range array {
		if x != 0 {
			return false
		}
	}
	return true
}

func process(input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		sequenceOfStrings := strings.Split(line, " ")
		sequence := make([]int, len(sequenceOfStrings))
		for i, s := range sequenceOfStrings {
			number, _ := strconv.Atoi(s)
			sequence[i] = number
		}
		total += processSequence(sequence)
	}

	return total
}

func main() {
	fmt.Println("Starting day09/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
