package main

import (
	"crossbow.de/aoc2023/util"
	_ "embed"
	"errors"
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

type Bag struct {
	r int
	g int
	b int
}

func (b *Bag) Add(o Bag) {
	if o.r > b.r {
		b.r = o.r
	}
	if o.g > b.g {
		b.g = o.g
	}
	if o.b > b.b {
		b.b = o.b
	}
}

func (b *Bag) String() string {
	return fmt.Sprintf("Bag(r: %d, g: %d, b: %d)", b.r, b.g, b.b)
}

func processAssignment(input string) Bag {
	trimmed := strings.Trim(input, " ")
	numAndColor := strings.Split(trimmed, " ")
	num, err := strconv.Atoi(strings.Trim(numAndColor[0], " "))
	if err != nil {
		panic(errors.New("could not convert string to int"))
	}
	color := numAndColor[1]

	switch color {
	case "red":
		return Bag{r: num}
	case "green":
		return Bag{g: num}
	case "blue":
		return Bag{b: num}
	}

	panic("unknown color")
}

func process(input string) int {
	lines := strings.Split(input, "\n")
	gameToBag := make(map[int]Bag)

	for _, line := range lines {
		titleAndSteps := strings.Split(line, ":")
		gameIndex, _ := strconv.Atoi(strings.Split(titleAndSteps[0], " ")[1])
		steps := strings.Split(titleAndSteps[1], ";")

		bag := Bag{}
		for _, step := range steps {
			assignments := strings.Split(step, ",")
			for _, assignment := range assignments {
				newBag := processAssignment(assignment)
				//fmt.Println(newBag)
				bag.Add(newBag)
			}
		}
		gameToBag[gameIndex] = bag
		//fmt.Println(bag)
	}

	sum := 0
	for _, bag := range gameToBag {
		sum += bag.r * bag.g * bag.b
	}

	return sum
}

func main() {
	fmt.Println("Starting day02/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
