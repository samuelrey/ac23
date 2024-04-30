package main

import (
	"ac23/day01"
	"ac23/day02"
	"ac23/day03"
	"ac23/day04"
	"ac23/day05"
	"bufio"
	"fmt"
	"log"
	"os"
)

type Puzzler interface {
	Part1([]string) string
	Part2([]string) string
}

func main() {
	puzzlerByID := map[string]Puzzler{
		"day01": day01.Puzzler{},
		"day02": day02.Puzzler{},
		"day03": day03.Puzzler{},
		"day04": day04.Puzzler{},
		"day05": day05.Puzzler{},
	}

	day := "day05"
	target := "input"

	puzzler := puzzlerByID[day]

	input := readInput(day, target)

	displayResult(day, "Part1", target, puzzler.Part1(input))
	// displayResult(day, "Part2", puzzler.Part2(input))
}

func readInput(day, target string) []string {
	filename := fmt.Sprintf("%s/%s", day, target)
	fmt.Printf("Reading %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func displayResult(day, part, target, result string) {
	fmt.Printf("%s/%s [%s]: %s\n", day, part, target, result)
}
