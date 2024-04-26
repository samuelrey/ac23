package day03

import (
	"fmt"
	"unicode"
)

type Puzzler struct {
	candidates map[string]Part
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

func (p *Puzzler) Part1(input []string) string {
	p.numRows = len(input)
	p.numCols = len(input[0])
	p.findPartCandidates(input)
	p.findSymbols(input)

	for _, part := range p.candidates {
		adjacentCoordinates := p.findAdjacentCoordinates(part)
		fmt.Println(adjacentCoordinates)
	}

	// Taking a break. The plan going forward is as follows:
	// * calculate valid, adjacent coordinates for each part candidate
	// * for each adjacent coordinate, check against the map of symbols
	// * if a symbol exists at an adjacent coordinate, the partID should be added to the sum
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func (p *Puzzler) findPartCandidates(input []string) map[string]Part {
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

func (p *Puzzler) findAdjacentCoordinates(part Part) []Coordinate {
	coordinates := []Coordinate{}

	colAhead := part.Location.Col - 1
	colBehind := part.Location.Col + len(part.ID)

	if coordinate := p.ensureIsValidCoordinate(part.Location.Row, colAhead); coordinate != nil {
		coordinates = append(coordinates, *coordinate)
	}

	if coordinate := p.ensureIsValidCoordinate(part.Location.Row, colBehind); coordinate != nil {
		coordinates = append(coordinates, *coordinate)
	}

	rowAhead := part.Location.Row - 1
	rowBehind := part.Location.Row + 1

	for _, row := range []int{rowAhead, rowBehind} {
		for col := colAhead; col <= colBehind; col++ {
			if coordinate := p.ensureIsValidCoordinate(row, col); coordinate != nil {
				coordinates = append(coordinates, *coordinate)
			}
		}
	}

	return coordinates
}

func (p *Puzzler) ensureIsValidCoordinate(row int, col int) *Coordinate {
	if row >= 0 && row < p.numRows && col >= 0 && col < p.numCols {
		return &Coordinate{row, col}
	}

	return nil
}
