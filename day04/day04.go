package day04

import (
	"fmt"
	"strings"
)

type Puzzler struct {
}

type Card map[string]struct{}

func (Puzzler) Part1(input []string) string {
	i := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	splitByID := strings.Split(i, ":")
	splitPullAndWinner := strings.Split(splitByID[1], "|")

	pull := buildCard(splitPullAndWinner[0])
	winner := buildCard(splitPullAndWinner[1])
	sum := sumPoints([]Card{pull}, []Card{winner})
	fmt.Println(sum)
	return splitPullAndWinner[0]
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
		if matches > 0 {
			points = points + 2 ^ (matches - 1)
		}
	}
	return points
}

func countMatches(pulls Card, winners Card) int {
	sum := 0
	for number := range pulls {
		_, match := winners[number]
		if match {
			sum = sum + 1
		}
	}
	return sum
}
