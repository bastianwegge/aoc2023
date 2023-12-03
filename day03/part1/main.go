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

func isDigit(char string) bool {
	return char >= "0" && char <= "9"
}

type Coordinate struct {
	X int
	Y int
}

func process(input string) int {
	// want
	startCoordinates := make(map[Coordinate]bool)
	sum := 0

	// parse input
	grid := make([][]string, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	// iterate over grid x / y and gather all starting points for numbers
	for y, line := range grid {
		for x, char := range line {
			// avoid . and numbers
			if char == "." || isDigit(char) {
				continue
			}

			// check adjacent points for numbers
			for _, currentY := range []int{y - 1, y, y + 1} {
				for _, currentX := range []int{x - 1, x, x + 1} {
					// skip out of bounds for x & y & numbers
					if currentY < 0 || currentY >= len(grid) ||
						currentX < 0 || currentX >= len(grid[currentY]) ||
						!isDigit(grid[currentY][currentX]) {
						continue
					}

					// we are left with hits on numbers only!
					// now we need to find the first digit and its x,y
					leftOuter := currentX
					for leftOuter > 0 && isDigit(grid[currentY][leftOuter-1]) {
						leftOuter--
					}

					// leftOuter is now the first digit
					startCoordinates[Coordinate{X: leftOuter, Y: currentY}] = true
				}
			}
		}
	}

	// iterate over starting points and find complete numbers
	for coordinate, _ := range startCoordinates {
		collector := ""
		for i := coordinate.X; i < len(grid[0]) && isDigit(grid[coordinate.Y][i]); i++ {
			collector += grid[coordinate.Y][i]
		}
		number, err := strconv.Atoi(collector)
		if err != nil {
			panic("could not convert string to int")
		}
		//debug:fmt.Println("c: ", collector)
		sum += number
	}

	return sum
}

func main() {
	fmt.Println("Starting day03/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
