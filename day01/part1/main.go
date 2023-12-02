package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func main() {
	input := readFile("input.txt")
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		firstNumber := 0
		lastNumber := 0

		for _, char := range line {
			character := string(char)
			if value, err := strconv.Atoi(character); err == nil {
				if firstNumber == 0 {
					firstNumber = value
				}
				lastNumber = value
			}
		}
		newNumber, err := strconv.Atoi(fmt.Sprintf("%d%d", firstNumber, lastNumber))
		if err != nil {
			log.Fatal(err)
		}
		sum += newNumber
	}

	fmt.Println(sum)
}
