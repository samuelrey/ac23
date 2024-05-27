package day10

import "fmt"

type Puzzler struct {
}

type Pipe struct {
	dir1 string
	dir2 string
}

type Position struct {
	row int
	col int
}

func findStart(input []string) Position {
	var start Position

	for row, line := range input {
		for col, r := range line {
			char := string(r)
			if char == "S" {
				start = Position{row, col}
			}
		}
	}

	return start
}

func findPipes(input []string) map[Position]Pipe {
	pipeByPosition := map[Position]Pipe{}
	pipeByID := map[string]Pipe{
		"|": {"N", "S"},
		"-": {"W", "E"},
		"F": {"S", "E"},
		"J": {"N", "W"},
		"L": {"N", "E"},
		"7": {"S", "W"},
	}

	for row, line := range input {
		for col, r := range line {
			pipe := string(r)
			position := Position{row, col}

			if pipe != "." && pipe != "S" {
				pipeByPosition[position] = pipeByID[pipe]
			}
		}
	}

	return pipeByPosition
}

func (Puzzler) Part1(input []string) string {
	start := findStart(input)
	pipes := findPipes(input)
	fmt.Println(start)
	fmt.Println(pipes)
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
