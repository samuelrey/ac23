package day02

import (
	"strconv"
	"strings"
)

type Puzzler struct {
}

func (Puzzler) Part1(input []string) string {
	gameIdSum := 0

	for _, i := range input {
		game := strings.Split(i, ":")
		draws := strings.Split(game[1], ";")
		validGame := true

		for _, d := range draws {
			if !isValidDraw(d) {
				validGame = false
			}
		}

		if validGame {
			gameIdStr := strings.Split(game[0], " ")[1]
			gameId, _ := strconv.Atoi(gameIdStr)
			gameIdSum = gameIdSum + gameId
		}
	}

	return strconv.Itoa(gameIdSum)
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func isValidDraw(i string) bool {
	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	cubes := strings.Split(i, ",")
	for _, c := range cubes {
		countAndColor := strings.Fields(c)
		count, _ := strconv.Atoi(countAndColor[0])
		color := countAndColor[1]

		if max[color] < count {
			return false
		}
	}

	return true
}
