package day04

type Puzzler struct {
}

type Card map[int]struct{}

type Winners map[int]struct{}

func (Puzzler) Part1(input []string) string {
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func sumPoints(cards []Card, winners []Winners) int {
	points := 0
	for i := range cards {
		matches := countMatches(cards[i], winners[i])
		if matches > 0 {
			points = points + 2 ^ (matches - 1)
		}
	}
	return points
}

func countMatches(card Card, winners Winners) int {
	sum := 0
	for number := range card {
		_, match := winners[number]
		if match {
			sum = sum + 1
		}
	}
	return sum
}
