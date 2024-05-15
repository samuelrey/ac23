package day08

import (
	"regexp"
)

type Puzzler struct {
}

type Direction struct {
	key   string
	left  string
	right string
}

func createDirection(s string) Direction {
	pattern := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)
	matches := pattern.FindStringSubmatch(s)
	return Direction{
		matches[1],
		matches[2],
		matches[3],
	}
}

func (Puzzler) Part1(input []string) string {
	traversalPattern := input[0]

	directionByKey := map[string]Direction{}
	for i := 2; i < len(input); i++ {
		d := createDirection(input[i])
		directionByKey[d.key] = d
	}

	currentPosition := "AAA"
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
