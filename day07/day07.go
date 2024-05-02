package day07

import (
	"strconv"
	"unicode"
)

type Puzzler struct {
}

func (Puzzler) Part1(input []string) string {
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func parseHand(hand string) []int {
	cardValues := make([]int, len(hand))

	for i, card := range hand {
		if unicode.IsDigit(card) {
			cardValue, _ := strconv.Atoi(string(card))
			cardValues[i] = cardValue
		} else {
			switch card {
			case 'T':
				cardValues[i] = 10
			case 'J':
				cardValues[i] = 11
			case 'Q':
				cardValues[i] = 12
			case 'K':
				cardValues[i] = 13
			case 'A':
				cardValues[i] = 14
			}
		}
	}

	return cardValues
}

/*
 it's easier to compare all numbers than figure out letters & numbers as strings
 2: 2
 3: 3
 ...
 T: 10
 J: 11
 Q: 12
 K: 13
 A: 14
*/
