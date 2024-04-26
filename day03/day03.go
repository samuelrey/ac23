package day03

import (
	"fmt"
	"strconv"
	"unicode"
)

type Puzzler struct {
	candidates map[Part]int
	symbols    map[Coordinate]string
	numRows    int
	numCols    int
}

type Part struct {
	ID       string
	Location Coordinate
}

type Coordinate struct {
	Row int
	Col int
}

func (p *Puzzler) Part1(input []string) string { // correct answer: 517021
	sum := 0

	p.numRows = len(input)
	p.numCols = len(input[0])
	p.findPartCandidates(input)
	p.findSymbols(input)

	for part := range p.candidates {
		sum = sum + p.findAdjacentSymbol(part)
	}

	return strconv.Itoa(sum)
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func (p *Puzzler) findPartCandidates(input []string) map[Part]int {
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
	}

	p.candidates = candidates
	return candidates
}

func (p *Puzzler) findSymbols(input []string) map[Coordinate]string {
	symbols := map[Coordinate]string{}

	for rowIndex, row := range input {
		for colIndex, rune := range row {
			if unicode.IsPunct(rune) || unicode.IsSymbol(rune) {
				if rune != '.' {
					newSymbol := Coordinate{rowIndex, colIndex}
					_, exists := symbols[newSymbol]
					if exists {
						fmt.Printf("Symbol at coordinates already exists: %v.\n", newSymbol)
					}
					symbols[Coordinate{rowIndex, colIndex}] = string(rune)
				}
			}
		}
	}

	p.symbols = symbols
	return symbols
}

func (p *Puzzler) findAdjacentSymbol(part Part) int {
	partID, err := strconv.Atoi(part.ID)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	colAhead := part.Location.Col - 1
	colBehind := part.Location.Col + len(part.ID)

	if coordinate := p.ensureIsValidCoordinate(part.Location.Row, colAhead); coordinate != nil {
		_, exists := p.symbols[*coordinate]
		if exists {
			return partID
		}
	}

	if coordinate := p.ensureIsValidCoordinate(part.Location.Row, colBehind); coordinate != nil {
		_, exists := p.symbols[*coordinate]
		if exists {
			return partID
		}
	}

	rowAhead := part.Location.Row - 1
	rowBehind := part.Location.Row + 1

	for _, row := range []int{rowAhead, rowBehind} {
		for col := colAhead; col <= colBehind; col++ {
			if coordinate := p.ensureIsValidCoordinate(row, col); coordinate != nil {
				_, exists := p.symbols[*coordinate]
				if exists {
					return partID
				}
			}
		}
	}

	return 0
}

func (p *Puzzler) ensureIsValidCoordinate(row int, col int) *Coordinate {
	if row >= 0 && row < p.numRows && col >= 0 && col < p.numCols {
		return &Coordinate{row, col}
	}

	return nil
}
