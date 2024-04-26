package day03

import (
	"strconv"
)

type Puzzler struct{}

func (p Puzzler) Part1(input []string) string {
	schematic := SchematicFromInput(input)
	sum := 0

	for part := range schematic.parts {
		sum = sum + schematic.findAdjacentSymbol(part)
	}

	return strconv.Itoa(sum)
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
