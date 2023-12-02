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

func calcSumForLinesWithNumbers(lines []string) int {
	numbersAsWords := []string{"one","two","three","four","five","six","seven","eight","nine"}
	sum := 0
	for _, line := range lines {
		digits := []int{}
		for i, char := range line {
			for index, number := range numbersAsWords {
				if strings.HasPrefix(line[i:], number) {
					digits = append(digits, index+1)
				}
			}
			character := string(char)
			if value, err := strconv.Atoi(character); err == nil {
				digits = append(digits, value)
			}
		}
		if len(digits) == 0 {
			break;
		}
		newNumber, err := strconv.Atoi(fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1]))
		fmt.Println(newNumber, digits)
		if err != nil {
			log.Fatal(err)
		}
		sum += newNumber
	}
	return sum
}

func main() {
	input := readFile("input2.txt")
	lines := strings.Split(input, "\n")
	sum := calcSumForLinesWithNumbers(lines)
	fmt.Println(sum)
}
