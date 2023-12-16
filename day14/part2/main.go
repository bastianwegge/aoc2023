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

func tiltNorth(grid [][]string) {
	// iterate and move pieces to the direction if they can move (no # above)
	for rowIndex, row := range grid {
		if rowIndex == 0 {
			continue
		}
		for colIndex, char := range row {
			if char == "O" {
				// walk to the top
				tmpIndex := rowIndex
				for tmpIndex > 0 && grid[tmpIndex-1][colIndex] == "." {
					grid[tmpIndex][colIndex] = "."
					grid[tmpIndex-1][colIndex] = "O"
					tmpIndex--
				}
			}
		}
	}
}

func transpose(slice [][]string) [][]string {
	xlen := len(slice[0])
	ylen := len(slice)
	result := make([][]string, xlen)
	for i := range result {
		result[i] = make([]string, ylen)
	}
	for i := 0; i < xlen; i++ {
		for j := 0; j < ylen; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func tiltEast(grid [][]string) {
	rowLen := len(grid[0])
	for colIndex := rowLen - 1; colIndex >= 0; colIndex-- {
		if colIndex == rowLen-1 {
			continue
		}

		for rowIndex, row := range grid {
			char := row[colIndex]
			if char == "O" {
				// walk to the top
				tmpCol := colIndex
				for tmpCol < len(row)-1 && grid[rowIndex][tmpCol+1] == "." {
					grid[rowIndex][tmpCol] = "."
					grid[rowIndex][tmpCol+1] = "O"
					tmpCol++
				}
			}
		}
	}
}

func GridToString(grid [][]string) string {
	result := make([]string, len(grid)*len(grid))
	for i, row := range grid {
		copy(result[i*len(grid):], row)
	}

	return strings.Join(result, "")
}

func process(input string) int {
	totalLoad := 0
	var grid [][]string
	for _, line := range strings.Split(input, "\n") {
		// add to grid
		gridLine := strings.Split(line, "")
		grid = append(grid, gridLine)
	}

	hashes := make(map[string]int)
	iterations := 1000000000
	for i := 0; i < iterations; i++ {
		tiltNorth(grid)

		// tilt west!
		grid = transpose(grid)
		tiltNorth(grid)
		grid = transpose(grid)

		// tilt south!
		grid = transpose(grid)
		tiltEast(grid)
		grid = transpose(grid)

		tiltEast(grid)

		hash := GridToString(grid)
		if _, exists := hashes[hash]; exists {
			// hash hit, we now know how long each cycle is.
			// Use mod len to fast-forward a lot of iterations
			i = iterations - (iterations-i)%(i-hashes[hash])
		}
		hashes[hash] = i
	}

	// count the load
	slices.Reverse(grid)
	for lineIndex, line := range grid {
		// count the number of O's
		oCount := strings.Count(strings.Join(line, ""), "O")
		totalLoad += oCount * (lineIndex + 1)
	}

	return totalLoad
}

func main() {
	fmt.Println("Starting day14/part1")
	inputText := util.ReadFile("input.txt")
	output := process(inputText)
	fmt.Println(output)
	return
}
