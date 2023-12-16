package day01

import (
	"regexp"
	"strconv"
)

type Puzzler struct {
}

func (Puzzler) Part1(input []string) string {
	pattern := regexp.MustCompile(`(\d)`)
	sum := 0
	for _, line := range input {
		numbers := pattern.FindAllString(line, -1)
		combinedFirstLast, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		sum = sum + combinedFirstLast
	}

	return strconv.Itoa(sum)
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
