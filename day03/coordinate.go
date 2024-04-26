package day03

type Coordinate struct {
	Row int
	Col int
}

func createValidCoordinate(row, col, maxRow, maxCol int) *Coordinate {
	if row >= 0 && row < maxRow && col >= 0 && col < maxCol {
		return &Coordinate{row, col}
	}

	return nil
}
