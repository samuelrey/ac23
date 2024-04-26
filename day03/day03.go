package day03

import (
	"strconv"
)

type Puzzler struct{}

func (p Puzzler) Part1(input []string) string {
	schematic := SchematicFromInput(input)
	sum := schematic.sumPartNumbers()
	return strconv.Itoa(sum)
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
