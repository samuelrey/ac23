package day03

import (
	"fmt"
	"strconv"
)

type Part struct {
	ID       string
	Location Coordinate
}

func (part Part) idIfAdjacentSymbolExists(symbols map[Coordinate]struct{}, numRows int, numCols int) int {
	partID, err := strconv.Atoi(part.ID)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	colAhead := part.Location.Col - 1
	colBehind := part.Location.Col + len(part.ID)

	if coordinate := createValidCoordinate(part.Location.Row, colAhead, numRows, numCols); coordinate != nil {
		_, exists := symbols[*coordinate]
		if exists {
			return partID
		}
	}

	if coordinate := createValidCoordinate(part.Location.Row, colBehind, numRows, numCols); coordinate != nil {
		_, exists := symbols[*coordinate]
		if exists {
			return partID
		}
	}

	rowAhead := part.Location.Row - 1
	rowBehind := part.Location.Row + 1

	for _, row := range []int{rowAhead, rowBehind} {
		for col := colAhead; col <= colBehind; col++ {
			if coordinate := createValidCoordinate(row, col, numRows, numCols); coordinate != nil {
				_, exists := symbols[*coordinate]
				if exists {
					return partID
				}
			}
		}
	}

	return 0
}
