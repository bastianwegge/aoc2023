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

type Rock struct {
	x int
	y int
}

func (rock *Rock) CanMoveNorth(grid [][]string) bool {
	if rock.y == 0 {
		return false
	}
	return grid[rock.y-1][rock.x] == "."
}

func process(input string) int {
	totalLoad := 0
	var grid [][]string
	var rocks []Rock
	for lineIndex, line := range strings.Split(input, "\n") {
		// add to grid
		gridLine := strings.Split(line, "")
		grid = append(grid, gridLine)
		// collect Rocks
		for columnIndex, char := range gridLine {
			if char == "O" {
				rocks = append(rocks, Rock{x: columnIndex, y: lineIndex})
			}
		}
	}

	// iterate and move pieces to the bottom if they can move (no # below)
	for _, rock := range rocks {
		for rock.CanMoveNorth(grid) {
			grid[rock.y][rock.x] = "."
			grid[rock.y-1][rock.x] = "O"
			rock.y--
		}
	}

	//fmt.Println("after movement")
	//for _, line := range grid {
	//	fmt.Println(line)
	//}

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
