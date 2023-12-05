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

type SourceToDestination struct {
	Source int
	Dest   int
	Count  int
}

type Conversion struct {
	kind         string
	SourceToDest []SourceToDestination
}

func (c *Conversion) GetForward(number int) int {
	minimum := -1
	for _, sourceToDest := range c.SourceToDest {
		if number >= sourceToDest.Source && number <= sourceToDest.Source+sourceToDest.Count {
			forward := number - sourceToDest.Source + sourceToDest.Dest
			if minimum == -1 || forward < minimum {
				minimum = forward
			}
		}
	}
	if minimum == -1 {
		return number
	}
	return minimum
}

func createRangeMap(input string) Conversion {
	lines := strings.Split(input, "\n")
	var sourceToDest []SourceToDestination

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		dest, _ := strconv.Atoi(numbers[0])
		source, _ := strconv.Atoi(numbers[1])
		count, _ := strconv.Atoi(numbers[2])
		sourceToDest = append(sourceToDest, SourceToDestination{Source: source, Dest: dest, Count: count})
	}

	return Conversion{SourceToDest: sourceToDest}
}

func process(input string) int {
	lines := strings.Split(input, "\n\n")
	seedsStarts := make([]int, 0)
	seedLengths := make([]int, 0)
	conversions := make([]Conversion, 0)

	for lineIndex, line := range lines {
		if lineIndex == 0 {
			seedStrings := strings.Split(line, " ")
			for seedIndex, seedString := range seedStrings {
				seed, err := strconv.Atoi(seedString)
				if err == nil {
					if seedIndex%2 == 0 {
						seedsStarts = append(seedsStarts, seed)
					} else {
						seedLengths = append(seedLengths, seed)
					}
				}
			}
			continue
		}

		titleAndConversion := strings.Split(line, ":\n")
		nextConversion := createRangeMap(titleAndConversion[1])
		nextConversion.kind = titleAndConversion[0]
		conversions = append(conversions, nextConversion)
	}
	if len(seedsStarts) != len(seedLengths) {
		panic("ahhhhhh 😱 seedsStarts and seedLengths not equal")
	}

	fmt.Println("created conversions")

	lowest := 0
	for _, seed := range seedsStarts {
		movingSeed := seed
		for _, conversion := range conversions {
			movingSeed = conversion.GetForward(movingSeed)
		}

		if lowest == 0 {
			lowest = movingSeed
		} else if movingSeed < lowest {
			lowest = movingSeed
		}
	}

	return lowest
}

func main() {
	fmt.Println("Starting day05/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
