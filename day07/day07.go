package day07

import (
	"sort"
	"strconv"
	"unicode"
)

type Puzzler struct {
}

type HandValue []int

type HandValues []HandValue

func (Puzzler) Part1(input []string) string {
	hands := []string{"32T3K", "T55J5", "KK677", "KTJJT", "QQQJA"}
	handValues := make(HandValues, len(hands))
	for i, hand := range hands {
		handValues[i] = parseCardValues(hand)
	}
	sort.Slice(handValues, handValues.sortHandsByValue)
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func parseCardValues(hand string) HandValue {
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

func calculateHandValue(hand []int) int {
	matches := map[int]int{}
	for _, card := range hand {
		count, exists := matches[card]
		if !exists {
			matches[card] = 1
		} else {
			matches[card] = count + 1
		}
	}

	// exit early if all cards are unique, ie High Card
	if len(matches) == 5 {
		return 0
	}

	// filter out single cards since they have no effect on hand value
	for card, count := range matches {
		if count == 1 {
			delete(matches, card)
		}
	}

	// either a Full-House or 2 Pairs
	if len(matches) == 2 {
		for _, count := range matches {
			if count == 3 {
				return 4
			}
		}

		return 2
	}

	// consider that Full-House > Three-of-a-Kind, 2 Pairs > Pair
	for _, count := range matches {
		switch count {
		case 5:
			return 6
		case 4:
			return 5
		case 3:
			return 3
		case 2:
			return 1
		}
	}

	return -1
}

func (h HandValues) sortHandsByValue(i, j int) bool {
	left, right := h[i], h[j]
	leftValue, rightValue := calculateHandValue(left), calculateHandValue(right)
	if leftValue == rightValue {
		for i := range left {
			if left[i] != right[i] {
				return left[i] < right[i]
			}
		}
	}

	return leftValue < rightValue
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
