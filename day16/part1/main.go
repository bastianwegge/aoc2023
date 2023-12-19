package main

import (
	"crossbow.de/aoc2023/util"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x int
	y int
}

func (p *Point) Move(direction string) Point {
	return Point{p.x + Directions[direction][1], p.y + Directions[direction][0]}
}

type Beam struct {
	Direction string
	Position  Point
}

// Directions where 0 is north, 1 is east, 2 is south, 3 is west
var Directions = map[string][]int{"^": {-1, 0}, ">": {0, 1}, "v": {1, 0}, "<": {0, -1}}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func isMovementOutOfBounds(grid []string, currentPosition Point, currentDirection string) bool {
	// if moving out of bounds, just break
	switch currentDirection {
	case "^":
		if currentPosition.y == 0 {
			return true
		}
	case ">":
		if currentPosition.x == len(grid[currentPosition.y])-1 {
			return true
		}
	case "v":
		if currentPosition.y == len(grid)-1 {
			return true
		}
	case "<":
		if currentPosition.x == 0 {
			return true
		}
	}
	return false
}

func getOppositeDirection(currentDirection string) string {
	switch currentDirection {
	case "^":
		return "v"
	case ">":
		return "<"
	case "v":
		return "^"
	case "<":
		return ">"
	default:
		panic("unknown direction")
	}
	return ""
}

func decideNextDirection(grid []string, currentPosition Point, currentDirection string) (string, bool) {
	switch grid[currentPosition.y][currentPosition.x] {
	case '.':
		return currentDirection, false
	case '|':
		if currentDirection == ">" || currentDirection == "<" {
			return "v", true
		}
		return currentDirection, false
	case '-':
		if currentDirection == "^" || currentDirection == "v" {
			return ">", true
		}
		return currentDirection, false
	case '/':
		if currentDirection == "v" {
			return "<", false
		}
		if currentDirection == "<" {
			return "v", false
		}
		if currentDirection == ">" {
			return "^", false
		}
		if currentDirection == "^" {
			return ">", false
		}
	case '\\':
		if currentDirection == "v" {
			return ">", false
		}
		if currentDirection == ">" {
			return "v", false
		}
		if currentDirection == "<" {
			return "^", false
		}
		if currentDirection == "^" {
			return "<", false
		}
	}
	fmt.Println("uncaught case", string(grid[currentPosition.y][currentPosition.x]), "at", currentPosition, "with direction", currentDirection)
	panic("unknown outcome for decideNextDirection")
	return "", false
}

func process(input string) int {
	grid := strings.Split(input, "\n")
	for _, row := range grid {
		fmt.Println(row)
	}

	beams := make([]Beam, 0)
	beams = append(beams, Beam{Direction: ">", Position: Point{0, 0}})
	beamOrigins := map[Point]string{}
	visitedPositions := map[Point]int{}

	// add current position to visited points
	visitedPositions[beams[0].Position] = 1
	for beamIndex := 0; beamIndex < len(beams); beamIndex++ {
		if beamIndex > 100000 {
			break
		}
		if beamIndex%10000 == 0 {
			fmt.Println("beam", beamIndex+1, "of", len(beams))
		}
		//fmt.Println("beam", beamIndex+1)
		beamIsProductive := true
		currentPosition := beams[beamIndex].Position
		currentDirection := beams[beamIndex].Direction
		slack := 500

		for beamIsProductive {
			// MOVING
			var splitBeam bool
			currentDirection, splitBeam = decideNextDirection(grid, currentPosition, currentDirection)

			// if moving out of bounds, just break
			if isMovementOutOfBounds(grid, currentPosition, currentDirection) {
				break
			}

			//fmt.Println("pos:", currentPosition, "dir:", currentDirection, "char:", string(grid[currentPosition.y][currentPosition.x]))
			if splitBeam {
				create := true
				if createdDirection, hasFoundCreatedBeam := beamOrigins[currentPosition]; hasFoundCreatedBeam {
					if createdDirection == currentDirection {
						create = false
					}
					fmt.Println("no need to spawn new beam, already exists")
				}
				if create {
					beams = append(beams, Beam{Direction: getOppositeDirection(currentDirection), Position: currentPosition})
				}
			}

			// move
			currentPosition = currentPosition.Move(currentDirection)

			// check if visited?
			if _, found := visitedPositions[currentPosition]; found {
				slack--
				if slack == 0 {
					//fmt.Println("slack is 0, breaking")
					beamIsProductive = false
				}
			} else {
				// register visitedPosition
				visitedPositions[currentPosition] = 1
			}
		}
	}

	//print visited in grid
	for y, row := range grid {
		for x, char := range row {
			if _, found := visitedPositions[Point{x, y}]; found {
				fmt.Print("X")
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}

	// count number of visitedBeams
	return len(visitedPositions)
}

func main() {
	fmt.Println("Starting day16/part1")
	inputText := util.ReadFile("input.txt")
	output := process(inputText)
	fmt.Println(output)
	return
}
