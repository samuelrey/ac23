package day03

import (
	"fmt"
	"unicode"
)

type Puzzler struct {
}

type Part struct {
	ID       string
	Location Coordinate
}

type Coordinate struct {
	Row int
	Col int
}

func (Puzzler) Part1(input []string) string {
	candidates := findPartCandidates(input)
	fmt.Println(candidates)
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func findPartCandidates(input []string) map[string]Part {
	candidates := map[string]Part{}

	for rowIndex, row := range input {
		isNewPart := true
		partID := ""
		partIndex := -1

		for index, rune := range row {
			if unicode.IsDigit(rune) {
				if isNewPart {
					partIndex = index
					partID = string(rune)
					isNewPart = !isNewPart
				} else {
					partID = partID + string(rune)
				}
			} else if len(partID) != 0 {
				_, exists := candidates[partID]
				if exists {
					fmt.Printf("PartID already exists in set of candidates: %s.\n", partID)
				}

				candidates[partID] = Part{
					partID,
					Coordinate{rowIndex, partIndex},
				}

				partID = ""
				partIndex = -1
			}
		}
	}

	return candidates
}
