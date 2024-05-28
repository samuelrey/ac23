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

	possiblePositions := []Position{
		{start.row - 1, start.col},
		{start.row + 1, start.col},
		{start.row, start.col - 1},
		{start.row, start.col + 1},
	}
	actualPositions := map[Position]Pipe{}
	for _, position := range possiblePositions {
		if _, exists := pipes[position]; exists {
			actualPositions[position] = pipes[position]
		}
	}

	stepByPositon := map[Position]int{}
	for position, pipe := range actualPositions {

	}
	return "Part1 not yet implemented."
}

func step(current Position, pipe Pipe, fromDirection string, pipeByPosition map[Position]Pipe) (Pipe, Position, string) {
	var stepDir, nextFrom string
	row, col := current.row, current.col

	if pipe.dir1 == fromDirection {
		stepDir = pipe.dir2
	} else {
		stepDir = pipe.dir1
	}

	switch stepDir {
	case "N":
		row--
		nextFrom = "S"
	case "S":
		row++
		nextFrom = "N"
	case "W":
		col--
		nextFrom = "E"
	case "E":
		col++
		nextFrom = "W"
	default:
		fmt.Println("d'oh")
	}

	nextPosition := Position{row, col}

	nextPipe, exists := pipeByPosition[nextPosition]
	if !exists {
		fmt.Println("uh oh")
	}

	return nextPipe, nextPosition, nextFrom
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}
