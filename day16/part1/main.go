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

func (p *Point) Move(direction int) Point {
	return Point{p.x + Directions[direction][0], p.y + Directions[direction][1]}
}

type Beam struct {
	Direction int
	Position  Point
}

// Directions where 0 is north, 1 is east, 2 is south, 3 is west
var Directions = map[int][]int{0: {-1, 0}, 1: {0, 1}, 2: {1, 0}, 3: {0, -1}}
var DirectionsDisplay = map[int]string{0: "^", 1: ">", 2: "v", 3: "<"}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func process(input string) int {
	grid := strings.Split(input, "\n")
	for _, row := range grid {
		fmt.Println(row)
	}

	beams := make([]Beam, 0)
	beams = append(beams, Beam{Direction: 1, Position: Point{0, 0}})
	visitedPositions := map[Point]int{}

	// add current position to visited points
	visitedPositions[beams[0].Position] = 1
	for _, beam := range beams {
		beamIsProductive := true
		currentPosition := beam.Position
		currentDirection := beam.Direction

		for beamIsProductive {
			// MOVING
			// if moving out of bounds, just break

			switch grid[currentPosition.y][currentPosition.x] {
			case '.':
				fmt.Println("moving", currentPosition, DirectionsDisplay[currentDirection])
				currentPosition = currentPosition.Move(currentDirection)
				fmt.Println("after moving", currentPosition)
			case '|':
				if currentDirection == 1 || currentDirection == 3 {
					fmt.Println("// add north beam, we go south")
					beams = append(beams, Beam{Direction: 2, Position: currentPosition})

					// check if movement would be out of bounds
					currentPosition = currentPosition.Move(currentDirection)
				} else {
					// just keep going
					currentPosition = currentPosition.Move(currentDirection)
				}
			default:
				break
			}

			// handle "|"
			// if next position is "|" and we're going east or west, split north and south
			// if next position is "|" and we're going north or south, keep going

			// handle "-"
			// if next position is "-" and we're going north or south, split east and west
			// if next position is "-" and we're going east or west, keep going

			// handle "/"
			// if next position is "/" and we're going south, go west
			// if next position is "/" and we're going west, go south
			// if next position is "/" and we're going east, go north
			// if next position is "/" and we're going north, go east

			// handle "\"
			// if next position is "\" and we're going south, go east
			// if next position is "\" and we're going east, go south
			// if next position is "\" and we're going west, go north
			// if next position is "\" and we're going north, go west

			// check if visited?
			fmt.Println(visitedPositions[currentPosition], currentPosition)
			if _, found := visitedPositions[currentPosition]; found {
				beamIsProductive = false
			} else {
				// register visitedPosition
				visitedPositions[currentPosition] = 1
			}
		}
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
