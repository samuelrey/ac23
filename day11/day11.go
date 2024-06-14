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

func (Puzzler) Part1(input []string) string {
	rowsToExpand, colsToExpand := []int{}, []int{}

	for i, row := range input {
		if !containsGalaxy(row) {
			rowsToExpand = append(rowsToExpand, i)
		}
	}

	transposed := transposeSpace(input)

	for i, row := range transposed {
		if !containsGalaxy(row) {
			colsToExpand = append(colsToExpand, i)
		}
	}

	prev := 0
	expanded := []string{}
	for _, index := range colsToExpand {
		expanded = append(expanded, transposed[prev:index]...)
		expanded = append(expanded, strings.Repeat(".", len(transposed[0])))
		prev = index
	}
	expanded = append(expanded, transposed[prev:]...)

	transposed = transposeSpace(expanded)
	prev = 0
	expanded = []string{}
	for _, index := range rowsToExpand {
		expanded = append(expanded, transposed[prev:index]...)
		expanded = append(expanded, strings.Repeat(".", len(transposed[0])))
		prev = index
	}
	expanded = append(expanded, transposed[prev:]...)
	
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
