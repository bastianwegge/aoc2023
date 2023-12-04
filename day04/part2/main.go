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

func process(input string) int {
	lines := strings.Split(input, "\n")
	cardWins := make([]int, 0)

	for cardIndex, line := range lines {
		titleAndNumbers := strings.Split(line, ":")
		numbersAndWinningNumbers := strings.Split(titleAndNumbers[1], "|")
		winningNumbers := strings.Split(numbersAndWinningNumbers[0], " ")
		numbers := strings.Split(numbersAndWinningNumbers[1], " ")

		cardWins = append(cardWins, 0)
		for _, winningNumber := range winningNumbers {
			if winningNumber == " " || winningNumber == "" {
				continue
			}
			for _, number := range numbers {
				if number == " " || number == "" {
					continue
				}

				if number == winningNumber {
					cardWins[cardIndex]++
				}
			}
		}
	}

	mapOfDuplication := make(map[int]int)
	for cardIndex := range lines {
		mapOfDuplication[cardIndex] = 1
	}
	for cardIndex, cardWinCount := range cardWins {
		factor := mapOfDuplication[cardIndex]
		for i := 1; i < cardWinCount+1; i++ {
			mapOfDuplication[cardIndex+i] += factor
		}
	}
	sum := 0
	for _, count := range mapOfDuplication {
		sum += count
	}
	return sum
}

func main() {
	fmt.Println("Starting day04/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
