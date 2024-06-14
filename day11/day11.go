package day11

import (
	"regexp"
	"strings"
)

type Puzzler struct {
}

func transposeSpace(space []string) []string {
	maxLength := len(space[0])

	result := make([]string, maxLength)
	for i := 0; i < maxLength; i++ {
		temp := ""
		for _, sector := range space {
			temp += string(sector[i])
		}
		result[i] = temp
	}

	return result
}

func containsGalaxy(sector string) bool {
	re := regexp.MustCompile(`[^.]`)
	return re.MatchString(sector)
}

func expandSpace(space []string) []string {
	expandIndexes := []int{}
	for i, sector := range space {
		if !containsGalaxy(sector) {
			expandIndexes = append(expandIndexes, i)
		}
	}

	prev := 0
	expanded := []string{}
	for _, index := range expandIndexes {
		expanded = append(expanded, space[prev:index]...)
		expanded = append(expanded, strings.Repeat(".", len(space[0])))
		prev = index
	}
	expanded = append(expanded, space[prev:]...)

	return expanded
}

func (Puzzler) Part1(space []string) string {
	space = expandSpace(space)
	space = transposeSpace(space)
	space = expandSpace(space)
	space = transposeSpace(space)
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
