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

func findAllGalaxy(sector string) []int {
	re := regexp.MustCompile(`#`)
	matches := re.FindAllStringIndex(sector, -1)
	indices := make([]int, len(matches))
	for i, match := range matches {
		indices[i] = match[0]
	}
	return indices
}

func expandSpace(space []string) []string {
	expandIndexes := []int{}
	for i, sector := range space {
		if len(findAllGalaxy(sector)) == 0 {
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

type galaxy struct {
	x int
	y int
}

func (Puzzler) Part1(space []string) string {
	space = expandSpace(space)
	space = transposeSpace(space)
	space = expandSpace(space)
	space = transposeSpace(space)

	galaxies := []galaxy{}
	for y, sector := range space {
		matches := findAllGalaxy(sector)
		for _, x := range matches {
			galaxies = append(galaxies, galaxy{x, y})
		}
	}

	galaxyPairs := [][]galaxy{}
	for i := 0; i < len(galaxies); i++ {
		for j := i+1; j < len(galaxies); j++ {
			galaxyPairs = append(galaxyPairs, []galaxy{galaxies[i], galaxies[j]})
		}
	}
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
