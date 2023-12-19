package main

import (
	"crossbow.de/aoc2023/util"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x int
	y int
}

// +Y direction
var North = Point{0, 1}

// -Y direction
var South = Point{0, -1}

// +X direction
var East = Point{1, 0}

// -X direction
var West = Point{-1, 0}

func (p *Point) Add(p2 Point) Point {
	return Point{p.x + p2.x, p.y + p2.y}
}

type Beam struct {
	Position  Point
	Direction Point
}

type Contraption struct {
	Tiles     map[Point]rune
	Energized map[Point]bool
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func NewContraption(grid []string) *Contraption {
	contraption := Contraption{make(map[Point]rune), make(map[Point]bool)}
	for i, row := range grid {
		for j, col := range row {
			contraption.Tiles[Point{j, -i}] = col
		}
	}
	return &contraption
}

func (c *Contraption) GetInitial() Beam {
	return Beam{Point{0, 0}, East}
}

func (c *Contraption) IsFinal(s Beam) bool {
	return false
}

func (c *Contraption) GetNeighbors(s Beam) []Beam {
	tile := c.Tiles[s.Position]

	switch tile {
	case 0:
		// off map
		return []Beam{}

	case '.':
		// continue on
		c.Energized[s.Position] = true
		newstate := Beam{s.Position.Add(s.Direction), s.Direction}
		return []Beam{newstate}

	case '\\':
		c.Energized[s.Position] = true
		var newDirection Point
		switch s.Direction {
		case North:
			newDirection = West
		case East:
			newDirection = South
		case South:
			newDirection = East
		case West:
			newDirection = North
		}
		newstate := Beam{s.Position.Add(newDirection), newDirection}
		return []Beam{newstate}

	case '/':
		c.Energized[s.Position] = true
		var newDirection Point
		switch s.Direction {
		case North:
			newDirection = East
		case East:
			newDirection = North
		case South:
			newDirection = West
		case West:
			newDirection = South
		}
		newstate := Beam{s.Position.Add(newDirection), newDirection}
		return []Beam{newstate}

	case '-':
		c.Energized[s.Position] = true
		switch s.Direction {
		case North, South:
			return []Beam{{s.Position.Add(East), East}, {s.Position.Add(West), West}}
		default:
			return []Beam{{s.Position.Add(s.Direction), s.Direction}}
		}

	case '|':
		c.Energized[s.Position] = true
		switch s.Direction {
		case East, West:
			return []Beam{{s.Position.Add(North), North}, {s.Position.Add(South), South}}
		default:
			return []Beam{{s.Position.Add(s.Direction), s.Direction}}
		}

	default:
		panic("Unexpected tile found")
	}
}

func process(input string) int {
	grid := strings.Split(input, "\n")
	contraption := NewContraption(grid)

	bfs := util.NewBFS[Beam]()

	_, err := bfs.Run(contraption)
	if err != util.BFSNotFound {
		log.Fatal(err, "Error before searching full map")
	}

	return len(contraption.Energized)
}

func main() {
	fmt.Println("Starting day16/part1")
	inputText := util.ReadFile("input.txt")
	output := process(inputText)
	fmt.Println(output)
	return
}
