package day03

import (
	"fmt"
	"strconv"
	"unicode"
)

type Schematic struct {
	parts   map[Part]int
	symbols map[Coordinate]string
	numRows int
	numCols int
}

func SchematicFromInput(input []string) Schematic {
	schematic := Schematic{}

	schematic.numRows = len(input)
	schematic.numCols = len(input[0])
	schematic.findPartCandidates(input)
	schematic.findSymbols(input)

	return schematic
}

func (s *Schematic) findPartCandidates(input []string) map[Part]int {
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

func (s *Schematic) findSymbols(input []string) map[Coordinate]string {
	symbols := map[Coordinate]string{}

	for rowIndex, row := range input {
		for colIndex, rune := range row {
			if !unicode.IsDigit(rune) && rune != '.' {
				newSymbol := Coordinate{rowIndex, colIndex}
				_, exists := symbols[newSymbol]
				if exists {
					fmt.Printf("Symbol at coordinates already exists: %v.\n", newSymbol)
				}
				symbols[Coordinate{rowIndex, colIndex}] = string(rune)
			}
		}
	}

	s.symbols = symbols
	return symbols
}

func (s *Schematic) findAdjacentSymbol(part Part) int {
	partID, err := strconv.Atoi(part.ID)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	colAhead := part.Location.Col - 1
	colBehind := part.Location.Col + len(part.ID)

	if coordinate := s.ensureIsValidCoordinate(part.Location.Row, colAhead); coordinate != nil {
		_, exists := s.symbols[*coordinate]
		if exists {
			return partID
		}
	}

	if coordinate := s.ensureIsValidCoordinate(part.Location.Row, colBehind); coordinate != nil {
		_, exists := s.symbols[*coordinate]
		if exists {
			return partID
		}
	}

	rowAhead := part.Location.Row - 1
	rowBehind := part.Location.Row + 1

	for _, row := range []int{rowAhead, rowBehind} {
		for col := colAhead; col <= colBehind; col++ {
			if coordinate := s.ensureIsValidCoordinate(row, col); coordinate != nil {
				_, exists := s.symbols[*coordinate]
				if exists {
					return partID
				}
			}
		}
	}

	return 0
}

func (s *Schematic) ensureIsValidCoordinate(row int, col int) *Coordinate {
	if row >= 0 && row < s.numRows && col >= 0 && col < s.numCols {
		return &Coordinate{row, col}
	}

	return nil
}
