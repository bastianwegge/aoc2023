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

type Race struct {
	time                int
	record              int
	waysToBeatTheRecord int
}

func (race *Race) fillWaysToBeatTheRecord() int {
	for i := 0; i < race.time; i++ {
		distance := 0
		speed := i
		leftTime := race.time - i
		distance = speed * leftTime
		if distance > race.record {
			race.waysToBeatTheRecord++
		}
	}
	return race.waysToBeatTheRecord
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func process(input string) int {
	lines := strings.Split(input, "\n\n")
	races := make([]Race, 0)
	// reading and parsing the input to Races
	for _, line := range lines {
		timeAndDistance := strings.Split(line, "\n")
		times := strings.Fields(timeAndDistance[0])
		distances := strings.Fields(timeAndDistance[1])
		for i, time := range times {
			if i == 0 {
				continue
			}
			timeAsInt, err := strconv.Atoi(time)
			if err != nil {
				panic(err)
			}
			distanceAsInt, err := strconv.Atoi(distances[i])
			if err != nil {
				panic(err)
			}

			race := Race{
				time:   timeAsInt,
				record: distanceAsInt,
			}
			races = append(races, race)
		}
	}

	sum := 1
	for _, race := range races {
		race.fillWaysToBeatTheRecord()
		sum *= race.waysToBeatTheRecord
	}
	return sum
}

func main() {
	fmt.Println("Starting day06/part2")
	input := util.ReadFile("input.txt")
	output := process(input)
	fmt.Println(output)
	return
}
