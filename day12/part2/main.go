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

var cache = make(map[string]int)

func findArrangements(springs string, groups []int) int {
	key := springs
	for _, group := range groups {
		key += strconv.Itoa(group) + ","
	}
	if v, ok := cache[key]; ok {
		return v
	}
	if len(springs) == 0 {
		if len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if strings.HasPrefix(springs, "?") {
		return findArrangements(strings.Replace(springs, "?", ".", 1), groups) +
			findArrangements(strings.Replace(springs, "?", "#", 1), groups)
	}
	if strings.HasPrefix(springs, ".") {
		res := findArrangements(strings.TrimPrefix(springs, "."), groups)
		cache[key] = res
		return res
	}

	if strings.HasPrefix(springs, "#") {
		if len(groups) == 0 {
			cache[key] = 0
			return 0
		}
		if len(springs) < groups[0] {
			cache[key] = 0
			return 0
		}
		if strings.Contains(springs[0:groups[0]], ".") {
			cache[key] = 0
			return 0
		}
		if len(groups) > 1 {
			if len(springs) < groups[0]+1 || string(springs[groups[0]]) == "#" {
				cache[key] = 0
				return 0
			}
			res := findArrangements(springs[groups[0]+1:], groups[1:])
			cache[key] = res
			return res
		} else {
			res := findArrangements(springs[groups[0]:], groups[1:])
			cache[key] = res
			return res
		}
	}

	return 0
}

func process(input string) int {
	total := 0
	for _, row := range strings.Split(input, "\n") {
		var comb []int
		configAndNumbers := strings.Split(row, " ")
		config := configAndNumbers[0]
		numbers := configAndNumbers[1]
		springs := ""
		for i := 0; i < 5; i++ {
			springs += config
			if i < 4 {
				springs += "?"
			}
		}
		for i := 0; i < 5; i++ {
			for _, number := range strings.Split(numbers, ",") {
				conv, _ := strconv.Atoi(number)
				comb = append(comb, conv)
			}
		}
		total += findArrangements(springs, comb)
	}

	return total
}

func main() {
	fmt.Println("Starting day12/part1")
	inputText := util.ReadFile("input.txt")
	output := process(inputText)
	fmt.Println(output)
	return
}
