package day04

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Puzzler struct {
}

type Card map[string]struct{}

func (Puzzler) Part1(input []string) string {
	sum := 0

	for _, row := range input {
		splitByID := strings.Split(row, ":")
		splitPullAndWinner := strings.Split(splitByID[1], "|")

		pull := buildCard(splitPullAndWinner[0])
		winner := buildCard(splitPullAndWinner[1])
		sum = sum + sumPoints([]Card{pull}, []Card{winner})
	}

	return strconv.Itoa(sum)
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func buildCard(s string) Card {
	card := Card{}
	numbers := strings.Fields(s)
	for _, n := range numbers {
		card[n] = struct{}{}
	}

	if len(numbers) != len(card) {
		fmt.Println("There are duplicates!!!")
	}

	return card
}

func sumPoints(pulls []Card, winners []Card) int {
	points := 0
	for i := range pulls {
		matches := countMatches(pulls[i], winners[i])
		points = points + calculatePointsForMatches(matches)
	}
	return points
}

func calculatePointsForMatches(matches int) int {
	if matches == 0 {
		return 0
	}

	return int(math.Pow(float64(2), float64(matches-1)))
}

func countMatches(pull Card, winner Card) int {
	sum := 0
	for number := range pull {
		_, match := winner[number]
		if match {
			sum = sum + 1
		}
	}
	return sum
}
