package day08

import (
	"fmt"
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

func traverse(directionByKey map[string]Direction, pattern string) int {
	count := 0
	step := 0
	currentPosition := "AAA"

	for currentPosition != "ZZZ" {
		d := directionByKey[currentPosition]
		r := pattern[step]

		switch r {
		case 'R':
			currentPosition = d.right
		case 'L':
			currentPosition = d.left
		default:
			panic(1)
		}

		count++
		step = (step + 1) % len(pattern)
	}

	return count
}

func (Puzzler) Part1(input []string) string {
	pattern := input[0]

	directionByKey := map[string]Direction{}
	for i := 2; i < len(input); i++ {
		d := createDirection(input[i])
		directionByKey[d.key] = d
	}

	x := traverse(directionByKey, pattern)
	fmt.Println(x)
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
