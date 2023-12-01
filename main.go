package main

import (
	"ac23/day01"
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
	day := "day01"
	target := "sample"

	puzzler := getPuzzler(day)

	input := readInput(day, target)

	displayResult(day, "Part1", puzzler.Part1(input))
	displayResult(day, "Part2", puzzler.Part2(input))
}

func getPuzzler(day string) Puzzler {
	var puzzler Puzzler
	switch day {
	case "day01":
		puzzler = day01.Puzzler{}
	default:
		log.Fatalf("%s: pkg does not exist, duplicate day00 to get started\n", day)
	}

	return puzzler
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

func displayResult(day string, part string, result string) {
	fmt.Printf("%s/%s: %s\n", day, part, result)
}
