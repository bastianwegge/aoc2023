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

type Point struct {
	x int
	y int
}

func findStartPoint(grid []string) Point {
	for y, row := range grid {
		for x, col := range row {
			if byte(col) == 'S' {
				return Point{x, y}
			}
		}
	}
	return Point{}
}

func findFarthestPointFrom(startPoint Point, grid []string) (int, map[Point]int) {
	visitedPoints := map[Point]int{startPoint: 0}
	uncheckedPoints := []Point{startPoint}

	farthestDistance := 0
	for len(uncheckedPoints) > 0 {
		current := uncheckedPoints[0]
		uncheckedPoints = uncheckedPoints[1:]
		connectedNeighbours := getConnectedNeighbours(grid, current)
		for _, point := range connectedNeighbours {
			if _, found := visitedPoints[point]; !found {
				visitedPoints[point] = visitedPoints[current] + 1
				farthestDistance = max(farthestDistance, visitedPoints[current]+1)
				uncheckedPoints = append(uncheckedPoints, point)
			}
		}
	}

	return farthestDistance, visitedPoints
}

func getConnectedNeighbours(grid []string, p Point) []Point {
	points := make([]Point, 0)
	switch grid[p.y][p.x] {
	case '|':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x, p.y - 1})
	case '-':
		points = append(points, Point{p.x + 1, p.y})
		points = append(points, Point{p.x - 1, p.y})
	case 'L':
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x + 1, p.y})
	case 'J':
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x - 1, p.y})
	case '7':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x - 1, p.y})
	case 'F':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x + 1, p.y})
	case '.':
	case 'S':
		down, right, up, left := grid[p.y+1][p.x], grid[p.y][p.x+1], grid[p.y-1][p.x], grid[p.y][p.x-1]
		if down == '|' || down == 'L' || down == 'J' {
			points = append(points, Point{p.x, p.y + 1})
		}
		if right == '-' || right == '7' || right == 'J' {
			points = append(points, Point{p.x + 1, p.y})
		}
		if up == '|' || up == '7' || up == 'F' {
			points = append(points, Point{p.x, p.y - 1})
		}
		if left == '-' || left == 'L' || left == 'F' {
			points = append(points, Point{p.x - 1, p.y})
		}
	}
	return points
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findStartTile(start Point, grid []string) byte {
	points := getConnectedNeighbours(grid, start)
	minx, maxx, miny, maxy := min(points[0].x, points[1].x), max(points[0].x, points[1].x), min(points[0].y, points[1].y), max(points[0].y, points[1].y)
	if points[0].x == points[1].x {
		return '|'
	} else if points[0].y == points[1].y {
		return '-'
	} else if minx < start.x && miny < start.y {
		return 'J'
	} else if maxx > start.x && maxy > start.y {
		return 'F'
	} else if maxx > start.x && miny < start.y {
		return 'L'
	} else if minx < start.x && maxy > start.y {
		return '7'
	}
	return '.'
}

func isInside(grid []string, p Point, theLoop map[Point]int) bool {
	if _, part := theLoop[p]; part {
		return false
	}
	count := 0
	cornerCounts := map[byte]int{}
	for y := p.y + 1; y < len(grid); y++ {
		check := Point{p.x, y}
		tile := grid[y][p.x]
		if tile == 'S' {
			tile = findStartTile(Point{p.x, y}, grid)
		}
		if _, part := theLoop[check]; part {
			if tile == '-' {
				count++
			} else if tile != '|' && tile != '.' {
				cornerCounts[tile]++
			}
		}
	}

	count += max(cornerCounts['L'], cornerCounts['7']) - abs(cornerCounts['L']-cornerCounts['7'])
	count += max(cornerCounts['F'], cornerCounts['J']) - abs(cornerCounts['F']-cornerCounts['J'])
	return count%2 == 1
}

func findEnclosedPoints(grid []string, visitedPoints map[Point]int) int {
	countInside := 0
	for y, row := range grid {
		for x := range row {
			if isInside(grid, Point{x, y}, visitedPoints) {
				countInside++
			}
		}
	}

	return countInside
}

func process(input string) int {
	grid := strings.Split(input, "\n")
	startPoint := findStartPoint(grid)
	_, visitedPoints := findFarthestPointFrom(startPoint, grid)
	return findEnclosedPoints(grid, visitedPoints)
}

func main() {
	fmt.Println("Starting day09/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
