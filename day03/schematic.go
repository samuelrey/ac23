package day03

import (
	"unicode"
)

type Schematic struct {
	parts   map[Part]struct{}
	symbols map[Coordinate]struct{}
	numRows int
	numCols int
}

func SchematicFromInput(input []string) Schematic {
	schematic := Schematic{}

	schematic.numRows = len(input)
	schematic.numCols = len(input[0])
	schematic.parts = findParts(input)
	schematic.symbols = findSymbols(input)

	return schematic
}

func findParts(input []string) map[Part]struct{} {
	parts := map[Part]struct{}{}

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
				coordinate := Coordinate{rowIndex, partIndex}
				part := Part{partID, coordinate}
				parts[part] = struct{}{}

				partID = ""
				partIndex = -1
				isNewPart = !isNewPart
			}
		}

		if len(partID) != 0 {
			coordinate := Coordinate{rowIndex, partIndex}
			part := Part{partID, coordinate}
			parts[part] = struct{}{}
		}
	}

	return parts
}

func findSymbols(input []string) map[Coordinate]struct{} {
	symbols := map[Coordinate]struct{}{}

	for rowIndex, row := range input {
		for colIndex, rune := range row {
			if !unicode.IsDigit(rune) && rune != '.' {
				coordinate := Coordinate{rowIndex, colIndex}
				symbols[coordinate] = struct{}{}
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
