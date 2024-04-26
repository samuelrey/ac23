package day03

import (
	"fmt"
	"unicode"
)

type Schematic struct {
	parts   map[Part]int
	symbols map[Coordinate]struct{}
	numRows int
	numCols int
}

func SchematicFromInput(input []string) Schematic {
	schematic := Schematic{}

	schematic.numRows = len(input)
	schematic.numCols = len(input[0])
	schematic.findParts(input)
	schematic.symbols = findSymbols(input)

	return schematic
}

func (s *Schematic) findParts(input []string) map[Part]int {
	candidates := map[Part]int{}

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
				candidate := Part{partID, Coordinate{rowIndex, partIndex}}
				_, exists := candidates[candidate]
				if exists {
					fmt.Printf("PartID already exists in set of candidates: %s.\n", partID)
				}

				candidates[candidate] = 0
				partID = ""
				partIndex = -1
				isNewPart = !isNewPart
			}
		}

		if len(partID) != 0 {
			candidate := Part{partID, Coordinate{rowIndex, partIndex}}
			_, exists := candidates[candidate]
			if exists {
				fmt.Printf("PartID already exists in set of candidates: %s.\n", partID)
			}

			candidates[candidate] = 0
		}
	}

	s.parts = candidates
	return candidates
}

func findSymbols(input []string) map[Coordinate]struct{} {
	symbols := map[Coordinate]struct{}{}

	for rowIndex, row := range input {
		for colIndex, rune := range row {
			if !unicode.IsDigit(rune) && rune != '.' {
				newSymbol := Coordinate{rowIndex, colIndex}
				_, exists := symbols[newSymbol]
				if exists {
					fmt.Printf("Symbol at coordinates already exists: %v.\n", newSymbol)
				}
				symbols[Coordinate{rowIndex, colIndex}] = struct{}{}
			}
		}
	}

	return symbols
}

func (s Schematic) sumPartNumbers() int {
	sum := 0

	for part := range s.parts {
		sum = sum + part.idIfAdjacentSymbolExists(s.symbols, s.numRows, s.numCols)
	}

	return sum
}
