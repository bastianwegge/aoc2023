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

type Point struct {
	x int
	y int
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

func findEmptyRows(grid []string) []int {
	emptyRows := make([]int, 0)
	for i, row := range grid {
		if strings.Count(row, ".") == len(row) {
			emptyRows = append(emptyRows, i)
		}
	}
	return emptyRows
}

func findEmptyCols(grid []string) []int {
	transposedGrid := transpose(grid)
	return findEmptyRows(transposedGrid)
}

func findGalaxies(grid []string) []Point {
	galaxies := make([]Point, 0)
	for i, row := range grid {
		for j, char := range row {
			if char == '#' {
				galaxies = append(galaxies, Point{x: j, y: i})
			}
		}
	}
	return galaxies
}

func process(input string) int {
	grid := strings.Split(input, "\n")

	emptyRows := findEmptyRows(grid)
	emptyCols := findEmptyCols(grid)
	galaxies := findGalaxies(grid)

	//fmt.Println(galaxies)
	//fmt.Println(emptyRows)
	//fmt.Println(emptyCols)
	//fmt.Println(grid)

	total := 0
	scale := 1000000
	for i, point1 := range galaxies {
		for _, point2 := range galaxies[:i] {
			for rowIndex := min(point1.y, point2.y); rowIndex < max(point1.y, point2.y); rowIndex++ {
				if slices.Contains(emptyRows, rowIndex) {
					total += scale
				} else {
					total += 1
				}
			}
			for colIndex := min(point1.x, point2.x); colIndex < max(point1.x, point2.x); colIndex++ {
				if slices.Contains(emptyCols, colIndex) {
					total += scale
				} else {
					total += 1
				}
			}
		}
	}

	return total
}

func main() {
	fmt.Println("Starting day10/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
