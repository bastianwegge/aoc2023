package main

import (
	"crossbow.de/aoc2023/util"
	_ "embed"
	"fmt"
	"slices"
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

type Lens struct {
	label string
	focal int
}

func (lens Lens) Equal(other Lens) bool {
	return lens.label == other.label
}

type Box []Lens

func (box *Box) add(lens Lens) {
	index := slices.IndexFunc(*box, lens.Equal)

	if index == -1 {
		*box = append(*box, lens)
		return
	}

	(*box)[index] = lens
}

func (box *Box) remove(lens Lens) {
	index := slices.IndexFunc(*box, lens.Equal)

	if index == -1 {
		return
	}

	*box = slices.Delete(*box, index, index+1)
}

func (box Box) power() int {
	var total int
	for i, lens := range box {
		total += (i + 1) * lens.focal
	}

	return total
}

type Boxes [256]Box

func (boxes *Boxes) processInstruction(instruction string) {
	switch instruction[len(instruction)-1] {
	case '-':
		label := instruction[:len(instruction)-1]
		hash := createHash(label)
		lens := Lens{label, 0}
		boxes[hash].remove(lens)
	default:
		label := instruction[:len(instruction)-2]
		hash := createHash(label)
		focal := int(instruction[len(instruction)-1] - '0')
		lens := Lens{label, focal}
		boxes[hash].add(lens)
	}
}

func createHash(str string) int {
	var currentHash int
	for _, char := range str {
		currentHash += int(char)
		currentHash *= 17
		currentHash %= 256
	}

	return currentHash
}

func process(input string) int {
	var boxes Boxes
	for _, instruction := range strings.Split(input, ",") {
		boxes.processInstruction(instruction)
		//hashSum += createHash(seq)
	}
	total := 0
	for i, box := range boxes {
		total += (i + 1) * box.power()
	}
	return total
}

func main() {
	fmt.Println("Starting day15/part2")
	inputText := util.ReadFile("input.txt")
	output := process(inputText)
	fmt.Println(output)
	return
}
