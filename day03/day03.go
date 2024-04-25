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
	symbols := findSymbols(input)
	fmt.Println(candidates)
	fmt.Println(symbols)
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

		for colIndex, rune := range row {
			if unicode.IsDigit(rune) {
				if isNewPart {
					partIndex = colIndex
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

func findSymbols(input []string) map[Coordinate]rune {
	symbols := map[Coordinate]rune{}

	for rowIndex, row := range input {
		for colIndex, rune := range row {
			if unicode.IsPunct(rune) || unicode.IsSymbol(rune) {
				if rune != '.' {
					newSymbol := Coordinate{rowIndex, colIndex}
					_, exists := symbols[newSymbol]
					if exists {
						fmt.Printf("Symbol at coordinates already exists: %v.\n", newSymbol)
					}
					symbols[Coordinate{rowIndex, colIndex}] = rune
				}
			}
		}
	}

	return symbols
}
