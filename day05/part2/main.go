package main

import (
	"crossbow.de/aoc2023/util"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"sync"
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
	//fmt.Println("forward", number, minimum)
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
						seedLengths = append(seedLengths, seed)
					} else {
						seedsStarts = append(seedsStarts, seed)
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

	var wg sync.WaitGroup
	lowest := make([]int, len(seedsStarts))
	for startsIndex, seed := range seedsStarts {
		wg.Add(1)
		fmt.Println("seed", seed)
		startsIndex := startsIndex
		seed := seed

		go func(index int, currentSeed int) {
			defer wg.Done()
			for i := 0; i < seedLengths[index]; i++ {
				movingSeed := currentSeed + i
				for _, conversion := range conversions {
					movingSeed = conversion.GetForward(movingSeed)
				}

				if lowest[index] == 0 {
					lowest[index] = movingSeed
				} else if movingSeed < lowest[index] {
					lowest[index] = movingSeed
				}
			}
		}(startsIndex, seed)
	}
	wg.Wait()

	minimum := lowest[0]
	for _, value := range lowest {
		if value < minimum {
			minimum = value
		}
	}

	return minimum
}

func main() {
	fmt.Println("Starting day05/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
