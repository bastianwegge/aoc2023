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
	points := 0
	for _, line := range lines {
		titleAndNumbers := strings.Split(line, ":")
		numbersAndWinningNumbers := strings.Split(titleAndNumbers[1], "|")
		winningNumbers := strings.Split(numbersAndWinningNumbers[0], " ")
		numbers := strings.Split(numbersAndWinningNumbers[1], " ")

		result := 0
		for _, winningNumber := range winningNumbers {
			if winningNumber == " " || winningNumber == "" {
				continue
			}
			for _, number := range numbers {
				if number == " " || number == "" {
					continue
				}

				if number == winningNumber {
					if result == 0 {
						result = 1
					} else {
						result = result * 2
					}
				}
			}
		}
		points += result
	}

	return points
}

func main() {
	fmt.Println("Starting day04/part1")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
