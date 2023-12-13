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

func transpose(grid []string) []string {
	transposed := make([]string, len(grid[0]))
	for _, row := range grid {
		for i, ch := range row {
			transposed[i] += string(ch)
		}
	}
	return transposed
}

func mirror(input []string) int {
	for i := 1; i < len(input); i++ {
		l := slices.Min([]int{i, len(input) - i})
		s1, s2 := slices.Clone(input[i-l:i]), input[i:i+l]
		slices.Reverse(s1)
		if smudge(s1, s2) {
			return i
		}
	}
	return 0
}

func smudge(s1, s2 []string) bool {
	diffs := 0
	for i := range s1 {
		for j := range s1[i] {
			if s1[i][j] != s2[i][j] {
				diffs++
			}
		}
	}
	return diffs == 1
}

func process(input string) int {
	total := 0
	puzzles := strings.Split(strings.TrimSpace(input), "\n\n")
	for _, puzzle := range puzzles {
		rows := make([]string, 0)
		for _, row := range strings.Fields(puzzle) {
			rows = append(rows, row)
		}
		cols := transpose(rows)
		total += mirror(cols) + 100*mirror(rows)
	}

	return total
}

func main() {
	fmt.Println("Starting day13/part1")
	inputText := util.ReadFile("input.txt")
	output := process(inputText)
	fmt.Println(output)
	return
}
